package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint16, p1 uint32, p2 int32) uint16 {
    p0 -= (p0 - (p0 * (^(p0))))
    p1 -= p1
    switch -812 {
    case 0:
        if false {
            p2 -= p2
        }
    default:
        var v4 uint = 0
        _ = v4
    }
    p0 = 0
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    switch 127 {
    case 42:
        for fi5 := 0; fi5 < 3; fi5++ {
            var v6 uint = 233
            switch "foo bar" {
            case "foo bar":
                for fi7 := 0; fi7 < 11; fi7++ {
                    v3[("a" + "foo bar")] = "foo bar"
                    _ = fi7
                }
            case "a":
                fmt.Println(v6)
            default:
                if v1 {
                    if true {
                        v8 := GenS1{F0: 2147483647}
                        v6 += v6
                        _ = v8
                    } else {
                        v9 := GenS2{F0: 255, F1: -425}
                        switch ("foo bar" + "foo bar") {
                        case "":
                            v1 = (false || v1)
                        default:
                            var v10 string = ("\n\t" + "a")
                            _ = v10
                        }
                        var v11 uint8 = 128
                        _ = v11
                        _ = v9
                    }
                }
            }
            fmt.Println(v1, v6)
            _ = fi5
            _ = v6
        }
    case 127:
        for fi12 := 0; fi12 < 4; fi12++ {
            if (1.5 > ((-2.25 / (-2.25 - 3.14159)) - 1.5)) {
                for fi13 := 0; fi13 < 20; fi13++ {
                    fmt.Println(fi13, fi13, fi13)
                    v1 = v1
                    v14 := []int{}
                    _ = fi13
                    _ = v14
                }
                fmt.Println(v1)
                v15 := GenS1{F0: 142}
                _ = v15
            }
            _ = fi12
        }
    }
    v16 := GenS2{F0: 42, F1: -365}
    v2[2147483647] = "\n\t"
    fmt.Println(v1)
    if ((32767 == 32767) && (!(-1 > -1))) {
        v1 = (true || (!true))
    }
    var v17 float64 = 0.0
    fmt.Println(v1)
    fmt.Println(v1, v17)
    var v18 int16 = 706
    v18 = -1
    var v19 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
