package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int64; F1 int64; F2 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 int8; F2 uint32 }
type GenS6 struct { F0 uint; F1 int32 }
func genFn0(p0 int32, p1 uint64) int {
    if false {
        v5 := GenS1{F0: ""}
        _ = v5
    } else {
        if (!(!true)) {
            var v6 float64 = 0.1
            _ = v6
        }
        v7 := []int{-513, 1, -576, 8}
        var v8 int16 = -1
        _ = v7
        _ = v8
    }
    var v9 uint64 = 42
    p0 -= p0
    var v10 int8 = -1
    for fi11 := 0; fi11 < 4; fi11++ {
        switch "foo bar" {
        case "hello":
            var v12 uint16 = uint16(v9)
            _ = v12
        }
        p0 = p0
        var v13 uint = 18446744073709551615
        _ = fi11
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 32767
}
func genFn1() int16 {
    var v14 float32 = (((1e10 - 0.1) / 1.0) * 0.1)
    var v15 uint16 = 262
    var v16 uint64 = 2
    if (!(0 <= 1)) {
        v17 := GenS1{F0: "a"}
        _ = v17
    } else {
        var v18 float32 = v14
        var v19 uint64 = (^(v16))
        v20 := []int{}
        _ = v18
        _ = v19
        _ = v20
    }
    if ((v14 * ((3.14159 + float32(1e10)) * (v14 * -2.25))) <= (v14 - (0.0 - (v14 * 0.0)))) {
        var v21 uint = 0
        v22 := []int{377, -625}
        _ = v21
        _ = v22
    } else {
        var v23 float64 = float64(float32(0.0))
        _ = v23
    }
    _ = v14
    _ = v15
    _ = v16
    return 42
}

func main() {
    var v1 int32 = -685
    var v2 int64 = 1
    v3 := []string{"\n\t", "a", "hello"}
    v4 := make(map[int]bool)
    v1 += v1
    fmt.Println(v1, v1)
    switch 333 {
    case -2147483648:
        v2 += (v2 % (((^(v2))) | 1))
    case 87:
        fmt.Println(v1)
    case 32767:
        var v24 int32 = v1
        _ = v24
    }
    v1 = -2147483648
    for fi25 := 0; fi25 < 6; fi25++ {
        if true {
            v1 += (v1 << ((65535) & 63))
            v26 := GenS2{F0: 65535, F1: 477}
            _ = v26
        } else {
            if true {
                var v27 float32 = 0.0
                _ = len(v3)
                _ = v27
            } else {
                var v28 int8 = -128
                if (32767 != 32767) {
                    var v29 int16 = -1
                    _ = v29
                }
                v30 := GenS1{F0: ("a" + "")}
                _ = v28
                _ = v30
            }
            v2 -= v2
            for fi31 := 0; fi31 < 2; fi31++ {
                if ((v2 + (-(v2))) >= (v2 & (-(v2)))) {
                    v32 := GenS3{F0: v2, F1: v2, F2: -134}
                    switch (fi31 >> ((18446744073709551615) & 63)) {
                    case 32767:
                        fmt.Println(fi31)
                    case 42:
                        var v33 uint = 0
                        _ = v33
                    default:
                        v2 = v2
                    }
                    var v34 float32 = ((0.1 / 1e10) * 0.0)
                    _ = v32
                    _ = v34
                } else {
                    var v35 string = ("\n\t" + ("foo bar" + (("\n\t" + "a") + ((("foo bar" + ("foo bar" + ("a" + "\n\t"))) + "\n\t") + "\n\t"))))
                    _ = v35
                }
                v36 := GenS1{F0: "foo bar"}
                v2 += v2
                _ = fi31
                _ = v36
            }
        }
        for fi37 := 0; fi37 < 11; fi37++ {
            var v38 uint = 426
            if len(v3) > 0 {
                v3[0] = ("hello" + "hello")
            }
            _ = fi37
            _ = v38
        }
        _ = fi25
    }
    if (v2 >= 949) {
        v39 := GenS2{F0: 128, F1: 42}
        if (1 >= 18446744073709551615) {
            var v40 int = 9223372036854775807
            _ = v40
        }
        for fi41 := 0; fi41 < 20; fi41++ {
            v42 := GenS1{F0: ("a" + "hello")}
            fmt.Println(v2)
            fmt.Println(fi41, fi41, v1)
            _ = fi41
            _ = v42
        }
        _ = v39
    }
    v43 := v4[127]
    var v44 int16 = 1000
    var v45 int = -2147483648
    v46 := GenS1{F0: "a"}
    v47 := GenS3{F0: (^(v2)), F1: v2, F2: (v44 << ((407) & 63))}
    if true {
        for fi48 := 0; fi48 < 3; fi48++ {
            if true {
                _ = len(v4)
            } else {
                v49 := GenS2{F0: 0, F1: 1}
                for fi50 := 0; fi50 < 14; fi50++ {
                    switch (fi50 + (^(fi48))) {
                    case 127:
                        v51 := GenS4{F0: ((("foo bar" + "hello") + ("a" + ("" + "hello"))) + "foo bar")}
                        _ = v51
                    case 42:
                        var v52 int16 = (v44 << ((2) & 63))
                        _ = v52
                    }
                    fmt.Println(v45, fi50, v2)
                    var v53 uint64 = 4294967295
                    _ = fi50
                    _ = v53
                }
                var v54 bool = true
                _ = v49
                _ = v54
            }
            _ = fi48
        }
        var v55 uint64 = 785
        fmt.Println(v55, v43)
        _ = v55
    }
    if (352 != 32767) {
        var v56 float32 = ((0.0 * (1e10 + 0.0)) * 1.5)
        _ = v56
    } else {
        v57 := GenS3{F0: 42, F1: (v2 - (v2 >> ((16) & 63))), F2: v44}
        v44 = (v44 | v44)
        if v43 {
            switch (v45 - genFn0(int32(v2), 18446744073709551615)) {
            case -9223372036854775808:
                v43 = false
            default:
                var v58 uint16 = 65535
                _ = v58
            }
            var v59 uint64 = 42
            fmt.Println(v44)
            _ = v59
        } else {
            if v43 {
                v60 := GenS5{F0: (v2 >> ((4294967295) & 63)), F1: 31, F2: 247}
                v1 -= 0
                _ = v60
            }
            var v61 uint64 = 0
            _ = v61
        }
        _ = v57
    }
    fmt.Println(v45, v43, v1)
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    switch "" {
    case "a":
        for fi62 := 0; fi62 < 5; fi62++ {
            var v63 int32 = v1
            _ = fi62
            _ = v63
        }
    case "foo bar":
        var v64 float64 = 0.1
        _ = v64
    case "\n\t":
        switch (("hello" + "") + "hello") {
        case "":
            var v65 uint16 = 65535
            _ = v65
        case "hello":
            var v66 int16 = (v44 & (v44 >> ((2) & 63)))
            _ = v66
        default:
            if v43 {
                if v43 {
                    _ = len(v4)
                    v45 = 69
                }
                v67 := GenS6{F0: 1, F1: v1}
                delete(v4, genFn0(v1, 0))
                _ = v67
            }
        }
    }
    if v43 {
        v68 := GenS2{F0: 643, F1: 255}
        v4[(v45 & (v45 / ((v45) | 1)))] = v43
        _ = v68
    } else {
        var v69 uint64 = 4294967295
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
