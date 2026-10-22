package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
func genFn0(p0 int32, p1 uint8) int64 {
    p1 = p1
    if (!true) {
        p1 += p1
        fmt.Println(p1, p1)
    } else {
        var v5 int16 = 32767
        var v6 float64 = 3.14159
        var v7 uint16 = 2
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 string = ((("\n\t" + "a") + ((("foo bar" + (("a" + "hello") + "a")) + "foo bar") + (("" + "a") + "a"))) + "hello")
    _ = p0
    _ = p1
    _ = v8
    return -1
}
func genFn1(p0 uint64, p1 int32, p2 int) string {
    var v9 uint = 4294967295
    if (true || true) {
        for fi10 := 0; fi10 < 9; fi10++ {
            for fi11 := 0; fi11 < 19; fi11++ {
                switch int(p1) {
                case 32768:
                    if (4294967295 > 0) {
                        v12 := []int{int(p1)}
                        var v13 int64 = 1
                        v13 -= (-(v13))
                        _ = v12
                        _ = v13
                    }
                case -2147483648:
                    v14 := GenS1{F0: (^(p2))}
                    _ = v14
                case 42:
                    var v15 float64 = 1.5
                    _ = v15
                }
                var v16 uint16 = 255
                v9 = (v9 / (((v9 / (((^(v9))) | 1))) | 1))
                _ = fi11
                _ = v16
            }
            _ = fi10
        }
        var v17 string = ""
        _ = v17
    } else {
        p0 += (p0 & 432)
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 uint8 = 2
            _ = fi18
            _ = v19
        }
        if (!false) {
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 bool = ((("hello" + "foo bar") + "hello") <= "")
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := GenS1{F0: (-(p2))}
                var v24 bool = true
                v24 = false
                _ = fi22
                _ = v23
                _ = v24
            }
        } else {
            v25 := []int{(-(p2)), (^(p2))}
            fmt.Println(p2)
            v26 := GenS2{F0: ("a" + ("hello" + "\n\t"))}
            _ = v25
            _ = v26
        }
    }
    if (!false) {
        if false {
            v27 := GenS2{F0: ""}
            fmt.Println(p1)
            _ = v27
        }
    }
    var v28 uint8 = 127
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v9
    return ""
}

func main() {
    var v1 int8 = 0
    v2 := []uint32{1, 2147483648, 0, 2147483647}
    var v3 uint = 18446744073709551615
    var v4 uint16 = 42
    v1 -= v1
    var v29 uint16 = (v4 + (v4 - (^(v4))))
    v30 := GenS2{F0: "hello"}
    for fi31 := 0; fi31 < 10; fi31++ {
        fmt.Println(v29, v1)
        v29 -= (v29 / ((v4) | 1))
        _ = fi31
    }
    v2 = append(v2, 650)
    v4 += (v4 % ((v29) | 1))
    if false {
        switch 4294967295 {
        case 56:
            v4 = v29
        }
        var v32 int32 = 33
        _ = v32
    }
    if len(v2) > 0 {
        v2[0] = 589
    }
    var v33 uint16 = v29
    v33 = 255
    var v34 float32 = 1.5
    v2 = append(v2, 1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
}
