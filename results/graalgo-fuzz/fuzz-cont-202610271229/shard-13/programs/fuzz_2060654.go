package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 bool }

func main() {
    v1 := []int{0, -2147483648, -764, -2147483648}
    var v2 string = "a"
    var v3 string = ((v2 + v2) + "hello")
    v4 := make(map[int]string)
    var v5 int64 = 0
    for fi6 := 0; fi6 < 11; fi6++ {
        if len(v1) > 0 {
            v1[0] = (^(fi6))
        }
        for fi7 := 0; fi7 < 12; fi7++ {
            fmt.Println(v3)
            v3 = v3
            var v8 int8 = 1
            _ = fi7
            _ = v8
        }
        v3 = ((v2 + v2) + ((v3 + v3) + v3))
        _ = fi6
    }
    v3 = v3
    fmt.Println(v2, v5)
    var v9 uint8 = 1
    for fi10 := 0; fi10 < 16; fi10++ {
        v2 = ((v2 + ((("a" + (v3 + "a")) + ((v3 + (("\n\t" + "hello") + v3)) + (v2 + v2))) + ("hello" + ("foo bar" + (v2 + v3))))) + "foo bar")
        if true {
            var v11 string = v2
            switch v3 {
            case "a":
                v5 = -1
            case "hello":
                var v12 uint = 0
                _ = v12
            default:
                delete(v4, (fi10 + 981))
            }
            if true {
                var v13 int = (fi10 | fi10)
                for fi14 := 0; fi14 < 5; fi14++ {
                    v11 = ((v2 + "\n\t") + "\n\t")
                    var v15 bool = ((v9 <= v9) && false)
                    v16 := GenS1{F0: 110}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                _ = v13
            }
            _ = v11
        }
        _ = fi10
    }
    v2 = "a"
    switch v3 {
    case "":
        if true {
            v3 = ("hello" + "")
            for fi17 := 0; fi17 < 19; fi17++ {
                if (0 >= v5) {
                    v18 := GenS1{F0: 127}
                    v5 += v5
                    for fi19 := 0; fi19 < 13; fi19++ {
                        v4[-569] = v2
                        var v20 uint64 = 0
                        _ = fi19
                        _ = v20
                    }
                    _ = v18
                } else {
                    v21 := GenS1{F0: -10}
                    var v22 int8 = -65
                    v2 = v3
                    _ = v21
                    _ = v22
                }
                var v23 uint32 = 1
                v3 = v3
                _ = fi17
                _ = v23
            }
        }
    }
    v24 := GenS2{F0: v3, F1: (!(!(729 <= 840)))}
    v4[-1] = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
