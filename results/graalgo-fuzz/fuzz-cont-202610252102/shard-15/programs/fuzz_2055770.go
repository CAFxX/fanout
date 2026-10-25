package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
func genFn0(p0 int16, p1 uint32) uint8 {
    if true {
        var v6 uint = 2
        fmt.Println(v6, v6, p0)
        _ = v6
    }
    fmt.Println(p1)
    var v7 float64 = 0.1
    if (true && false) {
        var v8 uint16 = 65535
        fmt.Println(v8, v7)
        v9 := GenS1{F0: (v8 <= 882)}
        _ = v8
        _ = v9
    } else {
        v10 := []int{-1}
        v7 = ((v7 / 1.0) * v7)
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    return 25
}
func genFn1() bool {
    var v11 uint32 = 0
    var v12 bool = ("\n\t" == "")
    v11 += v11
    v13 := []int{0}
    _ = v11
    _ = v12
    _ = v13
    return (-43 <= -128)
}

func main() {
    var v1 int8 = -1
    var v2 float32 = float32(1.5)
    var v3 bool = (true || true)
    v4 := []int{-290, 2147483647, 320}
    var v5 uint64 = 989
    var v14 string = (("foo bar" + "\n\t") + "\n\t")
    if v3 {
        if len(v4) > 0 {
            v4[0] = 0
        }
        switch ("hello" + (v14 + "\n\t")) {
        case "a":
            for fi15 := 0; fi15 < 6; fi15++ {
                if (("a" + v14) < v14) {
                    for fi16 := 0; fi16 < 1; fi16++ {
                        fmt.Println(v5, v2)
                        _ = fi16
                    }
                    var v17 bool = v3
                    var v18 int = fi15
                    _ = v17
                    _ = v18
                } else {
                    if (false && true) {
                        v1 = v1
                        var v19 int = (fi15 ^ (fi15 * 1))
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v19
                    }
                    var v20 uint = 18446744073709551615
                    _ = v20
                }
                var v21 int16 = 533
                fmt.Println(v5, v14, v14)
                _ = fi15
                _ = v21
            }
        }
    }
    fmt.Println(v2, v5, v5)
    v3 = (!(float32(1e10) < (v2 * 3.14159)))
    v3 = v3
    var v22 int16 = 0
    var v23 string = "\n\t"
    if false {
        fmt.Println(v3, v1)
    } else {
        var v24 uint64 = 1
        v3 = (!v3)
        _ = v24
    }
    v1 -= 42
    v25 := GenS2{F0: 18446744073709551615}
    v2 += (v2 - (0.1 / v2))
    var v26 bool = v3
    v27 := GenS2{F0: 558}
    var v28 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
