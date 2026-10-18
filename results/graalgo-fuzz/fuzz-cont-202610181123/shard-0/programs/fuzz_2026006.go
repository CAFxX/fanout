package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 bool }
type GenS2 struct { F0 uint16; F1 uint64; F2 bool }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint32; F1 uint64; F2 string }
type GenS5 struct { F0 int; F1 int; F2 int8 }
func genFn0(p0 uint8, p1 uint16) int16 {
    v4 := []int{-626, -237}
    fmt.Println(p1, p0)
    v5 := GenS1{F0: (((4294967295 != 875) && true) || false), F1: 1000, F2: (-2.25 != (3.14159 * 1e10))}
    v6 := GenS2{F0: 32767, F1: 4294967295, F2: ((false && (float32(1.5) <= (3.14159 + 1.5))) || (-357 <= int16(p0)))}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 32767
    v3 := make(map[int]string)
    if true {
        var v7 int8 = 1
        if false {
            v7 += -128
            var v8 bool = false
            fmt.Println(v8)
            _ = v8
        } else {
            v9 := v3[-917]
            if false {
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 uint = 18446744073709551615
                    var v12 int8 = (v7 | int8(fi10))
                    v12 -= v7
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                switch (v9 + (v9 + "foo bar")) {
                case "a":
                    var v13 uint32 = 2
                    _ = v13
                default:
                    fmt.Println(v7, v7, v2)
                }
            } else {
                fmt.Println(v7, v9)
                for fi14 := 0; fi14 < 4; fi14++ {
                    v9 = "hello"
                    v15 := []int{fi14, fi14, (fi14 | fi14), (fi14 - int(v7))}
                    v2 = v2
                    _ = fi14
                    _ = v15
                }
            }
            _ = v9
        }
        var v16 uint = 65535
        _ = v16
        _ = v7
    }
    v2 += 0
    if false {
        v17 := GenS2{F0: (v2 * v2), F1: 0, F2: true}
        v18 := GenS2{F0: 1, F1: 945, F2: false}
        _ = v17
        _ = v18
    } else {
        v19 := []int{0, 9223372036854775807, -158, 2147483647}
        for fi20 := 0; fi20 < 9; fi20++ {
            for fi21 := 0; fi21 < 9; fi21++ {
                v2 = (^(v2))
                if false {
                    _ = len(v1)
                } else {
                    var v22 int = 270
                    _ = v22
                }
                _ = fi21
            }
            v2 += v2
            v23 := v3[(fi20 | 358)]
            _ = fi20
            _ = v23
        }
        for fi24 := 0; fi24 < 12; fi24++ {
            for fi25 := 0; fi25 < 11; fi25++ {
                v2 = (v2 - (^(v2)))
                v26 := v1[(fi24 << ((528) & 63))]
                _ = fi25
                _ = v26
            }
            _ = fi24
        }
        _ = v19
    }
    if true {
        if true {
            var v27 int8 = 42
            _ = v27
        }
        v28 := GenS2{F0: 41, F1: 963, F2: true}
        _ = v28
    } else {
        v2 = (^(v2))
        for fi29 := 0; fi29 < 4; fi29++ {
            if false {
                var v30 uint = 1
                v2 -= (v2 % (((v2 | 57)) | 1))
                _ = v30
            } else {
                if ((!((!(false || (false && (fi29 == fi29)))) && (42 < 2))) && (false || ((true || (1 > 127)) && false))) {
                    var v31 uint = 0
                    _ = v31
                }
                var v32 uint8 = 0
                v33 := []int{}
                _ = v32
                _ = v33
            }
            var v34 uint16 = v2
            v35 := GenS3{F0: (fi29 << ((42) & 63)), F1: "a"}
            _ = fi29
            _ = v34
            _ = v35
        }
        if true {
            var v36 int = -1
            switch (("hello" + "hello") + "foo bar") {
            case "\n\t":
                if (v36 >= v36) {
                    var v37 uint = 171
                    var v38 float64 = 1.5
                    v39 := []int{(-(v36)), (v36 + (^(v36))), (v36 >> (((v37 << ((v37) & 63))) & 63))}
                    _ = v37
                    _ = v38
                    _ = v39
                } else {
                    var v40 string = "foo bar"
                    v2 = (v2 >> ((uint(v2)) & 63))
                    _ = v40
                }
            }
            _ = v36
        } else {
            var v41 uint64 = 4294967295
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        v43 := GenS4{F0: 409, F1: 212, F2: "foo bar"}
        _ = fi42
        _ = v43
    }
    if (true || (uint8(v2) >= 42)) {
        switch 0 {
        case 32768:
            for fi44 := 0; fi44 < 8; fi44++ {
                var v45 uint64 = 18446744073709551615
                switch 128 {
                case 127:
                    for fi46 := 0; fi46 < 19; fi46++ {
                        switch fi44 {
                        case 0:
                            var v47 bool = true
                            v48 := GenS4{F0: 716, F1: 4294967295, F2: ""}
                            v49 := []int{(fi44 * fi44), fi46, (-(fi46))}
                            if (!(v47 || (v47 && (!(true && v47))))) {
                                fmt.Println(fi46)
                                var v50 uint64 = v45
                                _ = v50
                            } else {
                                if (4294967295 != 499) {
                                    delete(v3, 0)
                                }
                                v51 := GenS4{F0: 1, F1: (v45 * v45), F2: "hello"}
                                _ = v51
                            }
                            _ = v47
                            _ = v48
                            _ = v49
                        }
                        _ = fi46
                    }
                }
                _ = fi44
                _ = v45
            }
        case -1:
            v2 = (^(v2))
        default:
            v52 := GenS2{F0: v2, F1: 18446744073709551615, F2: (!true)}
            _ = v52
        }
        v3[581] = "\n\t"
    }
    v2 = v2
    for fi53 := 0; fi53 < 4; fi53++ {
        if (fi53 != fi53) {
            for fi54 := 0; fi54 < 19; fi54++ {
                if (65535 == 0) {
                    v3[-2147483648] = ""
                }
                _ = fi54
            }
            fmt.Println(fi53, fi53)
        } else {
            v1[(-(fi53))] = 6
        }
        v55 := GenS5{F0: fi53, F1: 2147483647, F2: 0}
        v56 := GenS3{F0: fi53, F1: "hello"}
        _ = fi53
        _ = v55
        _ = v56
    }
    var v57 uint = 86
    v57 = v57
    v2 -= 65535
    if (("a" + "") >= (("a" + "\n\t") + ("a" + "hello"))) {
        if false {
            v2 = v2
            v58 := []int{-9223372036854775808, 2147483647, 127}
            if ((((53 != 1) || true) && (676 < 62)) && (uint64(v2) > uint64(v57))) {
                var v59 float64 = -2.25
                var v60 uint = (v57 << ((v57) & 63))
                var v61 float32 = 0.0
                _ = v59
                _ = v60
                _ = v61
            }
            _ = v58
        }
        v57 = v57
    }
    if (0 < int64(v2)) {
        switch 2147483647 {
        case 4294967295:
            v2 += (v2 >> ((321) & 63))
        default:
            v2 += (^(v2))
        }
    } else {
        var v62 uint32 = 2147483647
        fmt.Println(v57)
        _ = v62
    }
    _ = len(v3)
    v57 += 92
    if false {
        if ((1 != 0) || ((9223372036854775807 <= 9223372036854775807) || false)) {
            var v63 int = 0
            fmt.Println(v63, v57, v63)
            _ = v63
        }
        v1[2147483647] = 32767
        fmt.Println(v2, v57)
    }
    for fi64 := 0; fi64 < 15; fi64++ {
        var v65 int64 = -1
        _ = fi64
        _ = v65
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v57
}
