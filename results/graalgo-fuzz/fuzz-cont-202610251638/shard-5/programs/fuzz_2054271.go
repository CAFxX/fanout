package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint16, p1 string) int {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int = 32767
        v9 := []int{fi7, (fi7 - (v8 | v8)), fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        if false {
            p1 = ""
            for fi11 := 0; fi11 < 20; fi11++ {
                fmt.Println(fi10)
                p1 = p1
                _ = fi11
            }
        }
        p0 += (p0 / ((1) | 1))
        var v12 uint32 = 2147483648
        _ = fi10
        _ = v12
    }
    var v13 uint32 = 1
    _ = p0
    _ = p1
    _ = v13
    return 9223372036854775807
}
func genFn1(p0 uint, p1 int32) string {
    if (!true) {
        var v14 uint16 = 568
        _ = v14
    }
    if (false || false) {
        fmt.Println(p0)
        v15 := []int{32768, genFn0(0, "foo bar")}
        _ = v15
    } else {
        p0 -= 42
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 int32 = (p1 - (^(p1)))
            _ = fi16
            _ = v17
        }
        switch ("" + "a") {
        case "hello":
            var v18 string = "\n\t"
            _ = v18
        default:
            v19 := []int{}
            _ = v19
        }
    }
    v20 := GenS1{F0: 0, F1: 492}
    v21 := GenS2{F0: 15}
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return "\n\t"
}

func main() {
    var v1 uint16 = 0
    v2 := []string{"a", ("foo bar" + "foo bar"), ("hello" + ("\n\t" + ("foo bar" + "hello")))}
    var v3 int64 = -1
    var v4 int32 = 352
    v5 := make(map[int]uint32)
    v6 := make(map[string]uint32)
    if false {
        if (96 > 127) {
            delete(v5, 1)
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 float64 = 1e10
                if true {
                    v3 += (v3 << ((42) & 63))
                    var v24 int16 = 0
                    var v25 int32 = -1
                    _ = v24
                    _ = v25
                } else {
                    v23 = ((v23 - 0.1) + v23)
                }
                var v26 uint64 = 42
                _ = fi22
                _ = v23
                _ = v26
            }
        } else {
            for fi27 := 0; fi27 < 1; fi27++ {
                v28 := GenS1{F0: (^(v1)), F1: 1}
                _ = fi27
                _ = v28
            }
            v1 -= v1
        }
    }
    var v29 float64 = -2.25
    _ = len(v5)
    var v30 uint16 = (v1 - (v1 & (^(v1))))
    v1 += (v30 >> ((18446744073709551615) & 63))
    v31 := GenS1{F0: (^(v1)), F1: 421}
    v29 -= v29
    v32 := GenS1{F0: v1, F1: 2}
    for fi33 := 0; fi33 < 19; fi33++ {
        if (!(2147483648 == 329)) {
            v3 += (v3 + (v3 * -9223372036854775808))
        }
        var v34 uint8 = 255
        _ = fi33
        _ = v34
    }
    if false {
        var v35 int = 32767
        v29 -= v29
        if len(v2) > 0 {
            v2[0] = ("hello" + "hello")
        }
        _ = v35
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        _ = len(v2)
        _ = fi36
    }
    switch genFn0(372, "\n\t") {
    case 42:
        switch ("hello" + "a") {
        case "\n\t":
            for fi37 := 0; fi37 < 17; fi37++ {
                if false {
                    if false {
                        v5[int(v4)] = 59
                    }
                    var v38 int = 32768
                    v39 := GenS1{F0: (^(v30)), F1: 0}
                    _ = v38
                    _ = v39
                }
                v3 += v3
                _ = fi37
            }
        }
    case 32767:
        for fi40 := 0; fi40 < 2; fi40++ {
            switch "foo bar" {
            case "a":
                fmt.Println(v3, v3, v3)
            }
            _ = fi40
        }
    }
    v41 := GenS1{F0: v1, F1: 546}
    var v42 uint32 = 794
    for fi43 := 0; fi43 < 6; fi43++ {
        if true {
            switch ("" + "hello") {
            case "hello":
                var v44 string = ""
                _ = v44
            default:
                v1 -= 127
            }
            var v45 uint16 = (^(v1))
            _ = v45
        } else {
            delete(v6, "hello")
            if false {
                var v46 float32 = 3.14159
                var v47 float64 = v29
                v46 = 0.0
                _ = v46
                _ = v47
            }
            v4 = -1
        }
        _ = fi43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
