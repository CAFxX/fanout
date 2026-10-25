package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint, p1 int32, p2 uint) int {
    p1 += p1
    v6 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 446
}
func genFn1(p0 uint32, p1 uint) uint {
    p0 = (p0 % (((^(p0))) | 1))
    var v7 int16 = 1000
    _ = p0
    _ = p1
    _ = v7
    return (p1 - p1)
}

func main() {
    v1 := []string{(("a" + "\n\t") + ((("a" + "a") + "\n\t") + "\n\t")), "a"}
    v2 := make(map[int]string)
    var v3 float64 = 0.0
    var v4 int64 = -586
    v5 := []string{""}
    if false {
        v3 -= v3
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS1{F0: 42}
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 uint8 = 0
                v3 = float64(3.14159)
                _ = fi10
                _ = v11
            }
            var v12 uint16 = 1
            _ = fi8
            _ = v12
            _ = v9
        }
    }
    v4 += 0
    if ((42 > 127) || (false || false)) {
        _ = len(v1)
        v1 = append(v1, ("" + (("foo bar" + "foo bar") + ("\n\t" + "\n\t"))))
    } else {
        for fi13 := 0; fi13 < 17; fi13++ {
            fmt.Println(v3, v4)
            if true {
                v4 += v4
                fmt.Println(fi13)
                v4 -= v4
            }
            var v14 int16 = 32767
            _ = fi13
            _ = v14
        }
        if (false || (true && false)) {
            v4 += v4
            switch "a" {
            case "":
                if ((!(0 < 862)) || false) {
                    v4 -= (v4 & (v4 * -1))
                    var v15 uint32 = 202
                    _ = v15
                }
            }
        }
    }
    fmt.Println(v4, v4)
    delete(v2, -490)
    fmt.Println(v4, v4)
    var v16 string = ""
    var v17 uint16 = 1
    var v18 int32 = 181
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
