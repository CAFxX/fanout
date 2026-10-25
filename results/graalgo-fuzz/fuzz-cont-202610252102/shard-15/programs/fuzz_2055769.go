package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int16, p1 uint, p2 uint64) uint64 {
    for fi5 := 0; fi5 < 7; fi5++ {
        if true {
            p0 -= 1000
            var v6 int64 = 1
            fmt.Println(p1, fi5)
            _ = v6
        }
        p0 += -32768
        _ = fi5
    }
    v7 := GenS1{F0: 42}
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: 127}
        fmt.Println(p1, fi8, p1)
        fmt.Println(p0)
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 string = ""
        var v12 int16 = (p0 / (((p0 | p0)) | 1))
        _ = fi10
        _ = v11
        _ = v12
    }
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p2
}
func genFn1(p0 int, p1 int8, p2 uint) bool {
    p1 += -1
    v13 := GenS1{F0: 0}
    v14 := GenS1{F0: 253}
    v15 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return true
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = (true && (true && true))
    var v3 int64 = 0
    var v4 float64 = 1e10
    var v16 int = 647
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 bool = v2
        _ = fi17
        _ = v18
    }
    v19 := []int{}
    v3 = v3
    for fi20 := 0; fi20 < 11; fi20++ {
        for fi21 := 0; fi21 < 5; fi21++ {
            switch ("a" + (("foo bar" + (("\n\t" + "hello") + (("a" + "hello") + ((("hello" + "foo bar") + ("a" + "hello")) + "foo bar")))) + "")) {
            case "foo bar":
                if v2 {
                    fmt.Println(fi20, v16)
                } else {
                    v16 += fi21
                }
            default:
                var v22 float64 = v4
                _ = v22
            }
            _ = fi21
        }
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 bool = (42 > 2)
            _ = fi23
            _ = v24
        }
        _ = fi20
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        var v26 uint = 1
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        fmt.Println(v2, v16)
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        delete(v1, 65535)
        _ = fi27
    }
    delete(v1, 2147483647)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
