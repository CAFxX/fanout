package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[string]string)
    var v2 int8 = -128
    var v3 uint16 = 32767
    var v4 uint = 18446744073709551615
    var v5 int16 = -826
    fmt.Println(v5, v5)
    if false {
        _ = len(v1)
        v2 += (v2 >> ((v4) & 63))
    }
    var v6 int16 = (v5 - (-(v5)))
    var v7 int64 = -1
    v2 = v2
    if true {
        var v8 uint = (^(v4))
        v9 := GenS1{F0: (v4 >> (((v8 & 65535)) & 63))}
        v10 := GenS1{F0: 512}
        _ = v10
        _ = v8
        _ = v9
    } else {
        for fi11 := 0; fi11 < 4; fi11++ {
            switch fi11 {
            case 32767:
                fmt.Println(v5, v5, v6)
            case 32768:
                v5 = 32767
            case -66:
                v5 = (^(v5))
            default:
                for fi12 := 0; fi12 < 13; fi12++ {
                    var v13 uint32 = 512
                    if (true && false) {
                        v7 = (v7 | v7)
                    } else {
                        v3 = v3
                    }
                    var v14 string = "hello"
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
            v5 = (v5 << (((^(v4))) & 63))
            _ = fi11
        }
        if (false || (1 > -1)) {
            v4 = v4
            fmt.Println(v2)
            _ = len(v1)
        }
    }
    if false {
        fmt.Println(v3, v6)
        var v15 uint16 = v3
        var v16 int = int(v5)
        _ = v15
        _ = v16
    } else {
        _ = len(v1)
    }
    v3 = (v3 >> (((v4 * v4)) & 63))
    switch 32768 {
    case -10:
        var v17 int64 = v7
        _ = v17
    case -867:
        if (!(((v2 & v2) == (^(v2))) && (true || (v2 > v2)))) {
            for fi18 := 0; fi18 < 5; fi18++ {
                v3 = (v3 - v3)
                v19 := []int{}
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                _ = fi18
                _ = v19
            }
            v6 += -276
        } else {
            var v20 bool = (true && false)
            var v21 int8 = -105
            _ = v20
            _ = v21
        }
    case -748:
        fmt.Println(v2, v2)
    default:
        fmt.Println(v4, v6)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
