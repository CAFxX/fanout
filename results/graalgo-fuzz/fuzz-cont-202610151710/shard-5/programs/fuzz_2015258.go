package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint32, p1 uint8) uint8 {
    var v7 int32 = 0
    if ((275 > 9223372036854775807) && (false && (true || true))) {
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 uint32 = 640
            v7 = v7
            _ = fi8
            _ = v9
        }
        fmt.Println(v7, p1)
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        p0 -= (p0 - (p0 << ((892) & 63)))
        _ = fi10
    }
    var v11 uint64 = 4294967295
    fmt.Println(p0, p0, v7)
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return p1
}

func main() {
    v1 := []string{("a" + "hello"), ""}
    var v2 bool = (1e10 > 1.5)
    var v3 string = (("" + "a") + "a")
    v4 := []int{-9223372036854775808}
    v5 := make(map[uint32]bool)
    var v6 int16 = 960
    var v12 uint64 = 1
    fmt.Println(v2, v3)
    delete(v5, 1)
    if (!false) {
        for fi13 := 0; fi13 < 2; fi13++ {
            switch 128 {
            case -2147483648:
                if true {
                    var v14 int16 = (v6 - (v6 - 1))
                    var v15 int16 = (v6 + v6)
                    _ = v14
                    _ = v15
                }
            default:
                fmt.Println(v2, v6, v3)
            }
            if (!(1000 != v6)) {
                var v16 string = v3
                var v17 uint8 = 124
                switch fi13 {
                case -1:
                    v18 := GenS1{F0: fi13, F1: 101, F2: (v17 + (v17 * v17))}
                    _ = v18
                }
                _ = v16
                _ = v17
            }
            _ = fi13
        }
        v3 = ""
        v2 = (false || (!(v2 || (false && (v2 || (v2 && (v2 && (939 < 2))))))))
    } else {
        var v19 float64 = float64(1e10)
        v20 := GenS2{F0: 65535}
        _ = v19
        _ = v20
    }
    var v21 float64 = 0.1
    v6 -= (v6 & v6)
    var v22 int8 = 1
    v23 := v5[119]
    v6 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
