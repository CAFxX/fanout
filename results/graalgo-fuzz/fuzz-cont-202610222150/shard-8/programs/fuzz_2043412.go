package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 bool; F2 uint32 }
func genFn0(p0 int32, p1 int8, p2 uint) uint16 {
    for fi4 := 0; fi4 < 10; fi4++ {
        for fi5 := 0; fi5 < 17; fi5++ {
            v6 := GenS1{F0: 2}
            var v7 int64 = -9223372036854775808
            var v8 uint16 = 32767
            p0 -= 2147483647
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        v9 := []int{(fi4 % ((fi4) | 1))}
        if false {
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 float64 = ((1e10 - ((float64(float32(0.0)) / 0.1) - float64((0.0 * -2.25)))) / ((3.14159 * -2.25) + 3.14159))
                var v12 int64 = 0
                _ = fi10
                _ = v11
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 13; fi13++ {
                if len(v9) > 0 {
                    v9[0] = (fi4 | (fi4 - fi4))
                }
                p0 = p0
                v14 := GenS2{F0: p1, F1: false, F2: 2147483647}
                _ = fi13
                _ = v14
            }
            if false {
                fmt.Println(p2, p0, p2)
                for fi15 := 0; fi15 < 19; fi15++ {
                    if true {
                        var v16 bool = (float64(0.1) < 0.0)
                        p2 += (p2 / ((0) | 1))
                        p1 = (p1 * (p1 % ((p1) | 1)))
                        _ = v16
                    } else {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        p2 = (p2 - 0)
                    }
                    _ = len(v9)
                    var v17 bool = (!false)
                    _ = fi15
                    _ = v17
                }
                var v18 int32 = (p0 % ((p0) | 1))
                _ = v18
            } else {
                var v19 uint8 = 42
                _ = v19
            }
            var v20 int32 = p0
            _ = v20
        }
        _ = fi4
        _ = v9
    }
    p2 -= 787
    if (((("foo bar" + (("foo bar" + "hello") + "foo bar")) + "") + ("hello" + ("foo bar" + "a"))) != "") {
        var v21 int8 = 42
        var v22 float32 = float32((float64(1e10) + float64(float32(1.5))))
        if true {
            switch (("" + ("\n\t" + ((("foo bar" + (("foo bar" + "hello") + "\n\t")) + ((("a" + ("\n\t" + "\n\t")) + ("a" + ("foo bar" + "hello"))) + ("a" + ("" + ("hello" + ("hello" + "a")))))) + (("\n\t" + "") + "a")))) + "\n\t") {
            case "\n\t":
                fmt.Println(p1, p2, v22)
            default:
                v23 := []int{4294967295, 2147483647, 0, -1}
                _ = v23
            }
            for fi24 := 0; fi24 < 18; fi24++ {
                for fi25 := 0; fi25 < 14; fi25++ {
                    v26 := []int{}
                    _ = fi25
                    _ = v26
                }
                _ = fi24
            }
        } else {
            v27 := GenS1{F0: 2}
            _ = v27
        }
        _ = v21
        _ = v22
    }
    if true {
        var v28 uint64 = 18446744073709551615
        _ = v28
    }
    if false {
        p0 = (-(p0))
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 int8 = 42
    var v2 uint8 = 1
    v3 := []uint32{42, 565, 846, 2}
    v2 = (v2 / ((v2) | 1))
    fmt.Println(v2, v2)
    var v29 uint32 = 74
    var v30 float32 = 1e10
    var v31 string = "\n\t"
    v32 := GenS1{F0: v29}
    var v33 float64 = (float64(v30) / -2.25)
    var v34 uint64 = 18446744073709551615
    if false {
        if (v30 >= v30) {
            if len(v3) > 0 {
                v3[0] = v29
            }
        } else {
            fmt.Println(v34, v34, v33)
        }
        switch v31 {
        case "":
            v1 = 1
        default:
            if len(v3) > 0 {
                v3[0] = v29
            }
        }
    }
    v3 = append(v3, (v29 / ((v29) | 1)))
    var v35 int8 = (v1 - v1)
    for fi36 := 0; fi36 < 7; fi36++ {
        for fi37 := 0; fi37 < 5; fi37++ {
            var v38 int32 = 42
            _ = fi37
            _ = v38
        }
        v39 := GenS2{F0: -1, F1: true, F2: 42}
        _ = fi36
        _ = v39
    }
    v1 -= 52
    var v40 uint32 = v29
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v40
}
