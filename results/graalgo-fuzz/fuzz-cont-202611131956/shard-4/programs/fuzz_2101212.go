package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint32) int8 {
    if (735 < 4294967295) {
        v6 := []int{1}
        var v7 string = ((("hello" + "hello") + "") + "")
        if false {
            switch v7 {
            case "\n\t":
                v6 = append(v6, 42)
            default:
                var v8 uint32 = 1
                _ = v8
            }
            p0 += (^(p0))
        }
        _ = v6
        _ = v7
    } else {
        p0 = (p0 << ((189) & 63))
        v9 := GenS1{F0: -2147483648}
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        p0 = 65535
        _ = fi10
    }
    _ = p0
    return 100
}

func main() {
    v1 := []string{}
    v2 := []uint32{2, 42, 2, 2147483647}
    v3 := make(map[string]bool)
    v4 := []int{127, 127, 4294967295, -420}
    v5 := make(map[string]bool)
    var v11 int16 = 1
    var v12 int32 = -628
    var v13 int8 = -58
    var v14 float64 = 3.14159
    v15 := v5["foo bar"]
    v1 = append(v1, "a")
    switch 835 {
    case 127:
        for fi16 := 0; fi16 < 8; fi16++ {
            switch ("hello" + "hello") {
            case "\n\t":
                var v17 int64 = 1
                _ = v17
            case "foo bar":
                _ = len(v3)
            }
            v18 := GenS2{F0: 553}
            if true {
                switch "hello" {
                case "hello":
                    if (!(2147483647 >= 4294967295)) {
                        var v19 uint32 = 65535
                        v14 += 0.0
                        _ = v19
                    }
                case "foo bar":
                    for fi20 := 0; fi20 < 20; fi20++ {
                        var v21 int = fi20
                        _ = fi20
                        _ = v21
                    }
                default:
                    var v22 uint16 = 32767
                    _ = v22
                }
                for fi23 := 0; fi23 < 7; fi23++ {
                    v15 = (32767 < v11)
                    fmt.Println(v14, fi16, v15)
                    v24 := v3["hello"]
                    _ = fi23
                    _ = v24
                }
            }
            _ = fi16
            _ = v18
        }
    default:
        _ = len(v3)
    }
    var v25 int16 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
