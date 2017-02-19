require 'set'

class HPrinter
    def initialize
        @all = Set.new
    end

    def loadHierarchy(a, type = :root, b = a.class, seen = Set.new)
        key = [a,type,b]
        return if seen.include? key
        seen << key
        @all << key
        loadHierarchy(b, :is_a, b.class, seen) if b.respond_to? :class
        loadHierarchy(b, :extends, b.superclass, seen) if b.respond_to? :superclass
    end

    def printGraph
        puts "digraph {"
        @all.each do |lhs, edge, rhs|
            next if lhs.nil? || edge == :root
            rhs = 'nil' if rhs.nil?
            printf("\t%s -> %s [ label=\"%s\" color=\"%s\" ];\n", lhs.to_s, rhs.to_s, edge.to_s, edgeColor(edge))
        end
        puts "}"
    end

    def edgeColor(edge)
        case edge
        when :is_a
            'blue'
        when :extends
            'red'
        else
            'green'
        end
    end
end

hp = HPrinter.new

hp.loadHierarchy(1)
hp.loadHierarchy(true)
hp.loadHierarchy(false)
hp.loadHierarchy([])
hp.loadHierarchy('')
hp.loadHierarchy({})
hp.loadHierarchy(:a)
hp.loadHierarchy(Set.new)

hp.printGraph