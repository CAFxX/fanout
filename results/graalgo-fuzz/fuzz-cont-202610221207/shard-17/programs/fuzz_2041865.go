package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint64, p1 uint, p2 string) uint64 {
    fmt.Println(p1, p2)
    v4 := GenS1{F0: -1, F1: 81}
    v5 := []int{9223372036854775807}
    switch p2 {
    case "\n\t":
        var v6 uint = (^(p1))
        _ = v6
    default:
        p0 += p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1() int64 {
    var v7 int32 = -25
    var v8 bool = false
    for fi9 := 0; fi9 < 5; fi9++ {
        v7 = (v7 % ((2147483647) | 1))
        _ = fi9
    }
    v8 = false
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := []string{}
    v2 := []string{}
    var v3 uint32 = 873
    if ((-1 <= -32768) || true) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v3)
        v1 = append(v1, "hello")
    } else {
        switch 32768 {
        case 15:
            fmt.Println(v3)
        case 42:
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 float32 = (3.14159 + (3.14159 * 3.14159))
                _ = fi10
                _ = v11
            }
        default:
            if false {
                v3 -= v3
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3 += 0
    }
    v3 = (^(v3))
    for fi12 := 0; fi12 < 3; fi12++ {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        if (32767 < 255) {
            v1 = append(v1, "")
            var v13 string = ""
            if ((v3 >> ((42) & 63)) != (v3 | v3)) {
                for fi14 := 0; fi14 < 1; fi14++ {
                    var v15 float64 = 0.1
                    _ = fi14
                    _ = v15
                }
            } else {
                switch (fi12 & 32767) {
                case -2147483648:
                    fmt.Println(fi12)
                case 0:
                    for fi16 := 0; fi16 < 10; fi16++ {
                        if true {
                            v13 = "a"
                        }
                        var v17 uint8 = 255
                        switch fi12 {
                        case 0:
                            v13 = "hello"
                        }
                        _ = fi16
                        _ = v17
                    }
                }
            }
            _ = v13
        } else {
            v3 -= (v3 | 0)
            if (true || (((false && ((-9223372036854775808 >= genFn1()) || ((1e10 < 1e10) && ("a" != ("hello" + (("hello" + ((("foo bar" + "") + "") + "foo bar")) + ("a" + ""))))))) || true) && false)) {
                if len(v1) > 0 {
                    v1[0] = ""
                }
                var v18 uint8 = 128
                for fi19 := 0; fi19 < 11; fi19++ {
                    var v20 uint8 = (v18 << ((4294967295) & 63))
                    v20 += (v20 ^ v20)
                    var v21 bool = false
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = v18
            } else {
                v3 -= 42
                v3 -= v3
            }
            var v22 uint32 = 4294967295
            _ = v22
        }
        v3 = v3
        _ = fi12
    }
    var v23 int = 39
    if (!false) {
        var v24 uint32 = 80
        v25 := GenS1{F0: 42, F1: 9223372036854775807}
        _ = v24
        _ = v25
    } else {
        for fi26 := 0; fi26 < 17; fi26++ {
            v23 -= (fi26 << ((255) & 63))
            _ = fi26
        }
        var v27 uint = 1
        _ = v27
    }
    if false {
        v3 += 2147483648
        fmt.Println(v3)
        if len(v2) > 0 {
            v2[0] = ("foo bar" + "\n\t")
        }
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        v23 = fi28
        _ = fi28
    }
    _ = len(v1)
    fmt.Println(v23, v3)
    var v29 int64 = 42
    v23 -= v23
    for fi30 := 0; fi30 < 17; fi30++ {
        v2 = append(v2, (("\n\t" + "hello") + "\n\t"))
        var v31 int16 = 0
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        if (3.14159 >= 1e10) {
            v3 = (v3 * (v3 + (v3 << ((18446744073709551615) & 63))))
            var v33 uint32 = v3
            var v34 uint8 = 128
            _ = v33
            _ = v34
        } else {
            var v35 float32 = (3.14159 + ((1e10 + float32((0.0 + 1.5))) - 3.14159))
            _ = len(v2)
            _ = v35
        }
        if (!(868 <= fi32)) {
            v1 = append(v1, "a")
            v36 := GenS1{F0: -2147483648, F1: genFn1()}
            _ = v36
        }
        for fi37 := 0; fi37 < 6; fi37++ {
            v38 := GenS2{F0: 98}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi37
            _ = v38
        }
        _ = fi32
    }
    if ((!true) && (!false)) {
        var v39 uint32 = 61
        var v40 bool = true
        v39 = v39
        _ = v39
        _ = v40
    }
    var v41 int64 = genFn1()
    fmt.Println(v3, v41)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v41
}
