package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int16; F2 int }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint64, p1 uint64) uint8 {
    for fi6 := 0; fi6 < 1; fi6++ {
        p1 = 42
        _ = fi6
    }
    if false {
        p0 -= 1
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := []int{}
        if true {
            var v9 int16 = 42
            _ = v9
        } else {
            var v10 string = ("hello" + ("" + (("foo bar" + ("" + "a")) + ("" + ("foo bar" + ("" + "a"))))))
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    if ("foo bar" != ("a" + ("" + "hello"))) {
        if (!(!((255 >= 0) && true))) {
            p0 -= (p0 + p0)
            p0 = (p0 >> ((0) & 63))
            switch -2147483648 {
            case 1:
                for fi11 := 0; fi11 < 15; fi11++ {
                    p1 -= p1
                    for fi12 := 0; fi12 < 13; fi12++ {
                        p1 = (^(p0))
                        var v13 float32 = (((-2.25 - ((float32(0.0) * 1e10) + ((0.0 / -2.25) - 0.1))) - 3.14159) - (float32(1e10) / 0.1))
                        _ = fi12
                        _ = v13
                    }
                    v14 := GenS1{F0: "\n\t", F1: 1, F2: -93}
                    _ = fi11
                    _ = v14
                }
            case -2147483648:
                v15 := GenS1{F0: "foo bar", F1: 18446744073709551615, F2: 158}
                _ = v15
            }
        } else {
            p0 = p1
            var v16 uint64 = 0
            _ = v16
        }
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint8 = 2
    var v2 uint = 440
    var v3 float64 = (-2.25 / float64(0.1))
    var v4 float64 = 1.5
    var v5 int32 = -2147483648
    if false {
        v17 := GenS2{F0: 42}
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 bool = true
            _ = fi18
            _ = v19
        }
        v2 = 255
        _ = v17
    } else {
        for fi20 := 0; fi20 < 15; fi20++ {
            v21 := []int{(^(fi20))}
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
            v5 -= -577
            _ = fi20
            _ = v21
        }
        var v22 float64 = (v4 / v4)
        var v23 uint = 1
        _ = v22
        _ = v23
    }
    switch ("" + ("a" + (((("foo bar" + "") + "") + ((((("" + "foo bar") + "foo bar") + "") + (("\n\t" + "hello") + "hello")) + (((("\n\t" + "") + "") + "a") + "\n\t"))) + "a"))) {
    case "a":
        if true {
            var v24 string = "foo bar"
            v5 = v5
            v3 = ((v3 + float64((((1.5 - 0.0) * (1.5 / -2.25)) / float32(v3)))) / v4)
            _ = v24
        } else {
            v25 := GenS2{F0: 2}
            switch 129 {
            case -767:
                var v26 uint16 = 314
                _ = v26
            case 756:
                v27 := GenS2{F0: 850}
                _ = v27
            case -2147483648:
                for fi28 := 0; fi28 < 3; fi28++ {
                    v29 := []int{(fi28 | 32767), -447}
                    v2 = 492
                    _ = fi28
                    _ = v29
                }
            default:
                if false {
                    var v30 string = ((("\n\t" + "foo bar") + (("foo bar" + "foo bar") + "a")) + ("\n\t" + "hello"))
                    _ = v30
                }
            }
            _ = v25
        }
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        var v32 string = ("hello" + ("\n\t" + "\n\t"))
        v33 := GenS3{F0: "a", F1: 0, F2: -2147483648}
        for fi34 := 0; fi34 < 10; fi34++ {
            v35 := []int{int(v2), fi34, (fi31 << (((v2 | (v2 & v2))) & 63))}
            v3 += 1.5
            v32 = ""
            v36 := GenS1{F0: v32, F1: uint(v1), F2: fi31}
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    v4 -= float64((float32(v3) + (1e10 / 1.0)))
    if ("a" != "") {
        var v37 int = 4294967295
        switch ("a" + "a") {
        case "foo bar":
            if (-128 == -122) {
                switch (v37 * v37) {
                case 9223372036854775807:
                    v1 = 2
                case 1:
                    v38 := GenS2{F0: 4294967295}
                    _ = v38
                case -440:
                    if (!((1 >= -32768) || (v5 <= (-(v5))))) {
                        v39 := GenS3{F0: ((("a" + "\n\t") + "a") + "hello"), F1: 0, F2: 127}
                        v40 := GenS2{F0: 1}
                        _ = v39
                        _ = v40
                    }
                default:
                    v41 := []int{101, v37, (v37 - (^(v37)))}
                    _ = v41
                }
                var v42 int16 = -956
                _ = v42
            } else {
                v43 := GenS4{F0: 214}
                v44 := []int{}
                var v45 float64 = (-2.25 / (v4 * 1.5))
                _ = v43
                _ = v44
                _ = v45
            }
        case "hello":
            for fi46 := 0; fi46 < 20; fi46++ {
                v1 = (v1 % ((v1) | 1))
                _ = fi46
            }
        case "a":
            var v47 bool = (false && true)
            _ = v47
        default:
            v48 := GenS3{F0: "\n\t", F1: 0, F2: (v37 & (v37 % ((-616) | 1)))}
            _ = v48
        }
        _ = v37
    }
    v49 := GenS4{F0: 42}
    fmt.Println(v2, v1)
    fmt.Println(v5, v1)
    v50 := []int{32768}
    var v51 int64 = 42
    if true {
        v52 := GenS1{F0: "", F1: 0, F2: 622}
        _ = v52
    }
    v4 += float64(float32(v3))
    v53 := GenS4{F0: 146}
    v3 += ((0.1 / v3) / float64((((0.1 + 0.1) * 3.14159) / -2.25)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v53
}
