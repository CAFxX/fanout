package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint8; F1 uint; F2 uint16 }
func genFn0(p0 int16, p1 uint8) uint {
    for fi6 := 0; fi6 < 20; fi6++ {
        p0 += (p0 / (((p0 ^ (p0 * (p0 % (((p0 << ((4294967295) & 63))) | 1))))) | 1))
        v7 := GenS1{F0: (("a" + "\n\t") + "a")}
        p1 = (p1 | p1)
        _ = fi6
        _ = v7
    }
    if (!true) {
        var v8 int16 = -32768
        fmt.Println(p0, p0, v8)
        v9 := GenS1{F0: "a"}
        _ = v8
        _ = v9
    }
    var v10 int16 = (p0 & 42)
    if (0.0 >= 0.0) {
        if false {
            v11 := GenS2{F0: 181}
            for fi12 := 0; fi12 < 14; fi12++ {
                p0 = p0
                _ = fi12
            }
            p0 -= 847
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = v10
    return 65535
}
func genFn1(p0 bool, p1 int32, p2 int8) int8 {
    if true {
        if (621 >= 65535) {
            var v13 int32 = p1
            _ = v13
        } else {
            var v14 bool = false
            if p0 {
                v15 := GenS2{F0: 183}
                _ = v15
            }
            _ = v14
        }
    }
    switch ((("a" + "\n\t") + (("hello" + "\n\t") + "")) + ("" + (("\n\t" + (("\n\t" + "") + "\n\t")) + (("a" + (("hello" + ("\n\t" + "hello")) + "a")) + ("a" + "a"))))) {
    case "foo bar":
        switch 128 {
        case -240:
            fmt.Println(p2)
        case -9223372036854775808:
            var v16 int = 740
            _ = v16
        case 583:
            var v17 float32 = 0.0
            _ = v17
        }
    case "hello":
        var v18 float64 = 0.1
        _ = v18
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 int32 = 1
        p1 = p1
        v21 := []int{(fi19 / (((fi19 | (fi19 ^ 362))) | 1)), (fi19 - (fi19 / (((fi19 | fi19)) | 1))), 4294967295, (fi19 | (-(fi19)))}
        _ = fi19
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 float64 = float64((0.1 / 1.0))
    var v2 float32 = float32(1.5)
    v3 := make(map[uint32]uint32)
    var v4 int32 = -1
    v5 := make(map[int]bool)
    var v22 float64 = ((v1 * v1) * (1e10 - 1.5))
    delete(v3, 4294967295)
    v4 -= v4
    var v23 bool = (!false)
    var v24 int64 = -479
    for fi25 := 0; fi25 < 6; fi25++ {
        if (v22 >= v1) {
            for fi26 := 0; fi26 < 12; fi26++ {
                v27 := GenS1{F0: "\n\t"}
                v28 := GenS3{F0: 2}
                _ = fi26
                _ = v27
                _ = v28
            }
            for fi29 := 0; fi29 < 5; fi29++ {
                var v30 bool = v23
                v31 := GenS4{F0: 1, F1: (v24 + v24), F2: 0}
                if (false || ((v1 + v22) <= v22)) {
                    v4 = int32(fi29)
                    v30 = false
                }
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        _ = fi25
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        v24 -= -1
        _ = fi32
    }
    v2 -= 3.14159
    if true {
        v33 := GenS4{F0: -1, F1: -1, F2: 280}
        if (9223372036854775807 == (^(v24))) {
            v2 -= v2
            if (((140 > 42) && v23) && true) {
                fmt.Println(v24, v24, v22)
                fmt.Println(v4, v2, v1)
            }
            v3[188] = 613
        } else {
            v2 += (float32((v1 / v1)) * (v2 / (0.1 - (float32(v1) - (0.0 - (v2 * v2))))))
            switch ("foo bar" + ("foo bar" + "\n\t")) {
            case "":
                fmt.Println(v23)
            }
            v34 := GenS5{F0: 128, F1: 18446744073709551615, F2: 916}
            _ = v34
        }
        if v23 {
            v35 := []int{128, 0, 32768}
            v4 = (v4 >> ((1) & 63))
            _ = v35
        } else {
            var v36 uint8 = 42
            var v37 string = ("\n\t" + (((("foo bar" + ((("a" + "\n\t") + ((("\n\t" + ("foo bar" + "a")) + "hello") + "\n\t")) + "\n\t")) + "hello") + "\n\t") + ("foo bar" + "hello")))
            var v38 int64 = v24
            _ = v36
            _ = v37
            _ = v38
        }
        _ = v33
    }
    v39 := GenS4{F0: 0, F1: (v24 / ((v24) | 1)), F2: 8}
    switch -504 {
    case 4294967295:
        var v40 uint = 4294967295
        _ = v40
    case 2147483647:
        v41 := GenS1{F0: ""}
        _ = v41
    default:
        fmt.Println(v4)
    }
    for fi42 := 0; fi42 < 20; fi42++ {
        v23 = v23
        for fi43 := 0; fi43 < 7; fi43++ {
            var v44 int32 = -787
            v4 -= v44
            if v23 {
                var v45 float32 = float32(v22)
                if (((2147483648 != 4294967295) && (0 > (v24 ^ 42))) || ((-2147483648 == v4) && (-1 <= 31))) {
                    fmt.Println(v44)
                    if (712 <= 0) {
                        var v46 int32 = (^(v4))
                        _ = v46
                    }
                }
                v47 := []int{(fi43 & fi42), fi42, 627}
                _ = v45
                _ = v47
            }
            _ = fi43
            _ = v44
        }
        _ = len(v3)
        _ = fi42
    }
    for fi48 := 0; fi48 < 9; fi48++ {
        var v49 float64 = 1e10
        _ = fi48
        _ = v49
    }
    fmt.Println(v1, v22, v22)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
