package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 uint64; F1 uint; F2 uint8 }
func genFn0(p0 bool, p1 string) uint64 {
    p0 = (320 > 367)
    p1 = "hello"
    for fi6 := 0; fi6 < 17; fi6++ {
        p0 = (!p0)
        p0 = (143 <= 255)
        _ = fi6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        p0 = (!((!(!false)) || (!false)))
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    switch p1 {
    case "":
        var v9 bool = (0.1 >= -2.25)
        _ = v9
    }
    _ = p0
    _ = p1
    return 632
}
func genFn1(p0 uint32, p1 uint64, p2 uint32) int {
    var v10 uint16 = 852
    v11 := GenS1{F0: -734, F1: 640, F2: 692}
    var v12 uint64 = (p1 | (^(p1)))
    v10 = (^(v10))
    var v13 int16 = -288
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = -1
    var v3 uint64 = 2
    v4 := []string{"", "", "\n\t", "hello"}
    var v5 uint = 255
    v14 := GenS2{F0: 255}
    v5 = 42
    v15 := GenS1{F0: int32(v3), F1: 128, F2: 25}
    v5 -= v5
    var v16 uint32 = 13
    if (((v3 - genFn0((((("a" > ("" + "")) || false) || (!(3.14159 <= 1e10))) || (!((42 <= 1000) && false))), "\n\t")) != (v3 - (v3 * v3))) && false) {
        v3 = 18446744073709551615
    } else {
        var v17 uint16 = 243
        v18 := GenS3{F0: ("hello" + "\n\t")}
        _ = v17
        _ = v18
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (563 >= 65535) {
        v19 := GenS4{F0: v2, F1: "foo bar"}
        _ = v19
    } else {
        v4 = append(v4, "\n\t")
        if (-32768 == -74) {
            switch 32767 {
            case -2147483648:
                v4 = append(v4, "foo bar")
            default:
                for fi20 := 0; fi20 < 11; fi20++ {
                    v3 -= 429
                    fmt.Println(v2, fi20, fi20)
                    var v21 uint64 = v3
                    _ = fi20
                    _ = v21
                }
            }
        } else {
            v22 := GenS5{F0: 2, F1: (v5 | 2), F2: 1}
            fmt.Println(v3, v2, v2)
            _ = v22
        }
    }
    var v23 string = "hello"
    if (!true) {
        var v24 int16 = 42
        _ = v24
    }
    if false {
        switch 32767 {
        case -14:
            switch v23 {
            case "hello":
                for fi25 := 0; fi25 < 13; fi25++ {
                    for fi26 := 0; fi26 < 16; fi26++ {
                        v23 = v23
                        var v27 int16 = -32768
                        fmt.Println(v2, v5)
                        _ = fi26
                        _ = v27
                    }
                    if true {
                        var v28 uint32 = (v16 + 42)
                        _ = v28
                    }
                    v4 = append(v4, v23)
                    _ = fi25
                }
            }
        case -429:
            fmt.Println(v5, v3, v23)
        case 4294967295:
            var v29 uint16 = 528
            _ = v29
        }
        var v30 uint16 = 2
        _ = v30
    }
    v4 = append(v4, "\n\t")
    for fi31 := 0; fi31 < 7; fi31++ {
        var v32 bool = true
        if (v32 || (-1 != -2147483648)) {
            if (0 == 42) {
                if (258 != 1) {
                    for fi33 := 0; fi33 < 6; fi33++ {
                        var v34 int16 = -224
                        _ = fi33
                        _ = v34
                    }
                }
                v35 := v1[v16]
                var v36 uint64 = uint64(v5)
                _ = v35
                _ = v36
            }
            if ((!v32) || (1 <= (v3 + uint64(v5)))) {
                switch ("" + "foo bar") {
                case "":
                    fmt.Println(v23)
                case "foo bar":
                    v37 := GenS1{F0: 42, F1: fi31, F2: v5}
                    _ = v37
                }
                fmt.Println(v3)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            if len(v4) > 0 {
                v4[0] = v23
            }
        } else {
            if (fi31 == (fi31 * int(v16))) {
                var v38 int16 = 774
                var v39 int16 = 1000
                if v32 {
                    v40 := GenS2{F0: 0}
                    if true {
                        var v41 int = genFn1(v16, (v3 << ((v5) & 63)), (v16 << ((v5) & 63)))
                        _ = v41
                    }
                    fmt.Println(v39, fi31, v39)
                    _ = v40
                }
                _ = v38
                _ = v39
            } else {
                var v42 int = fi31
                v43 := GenS2{F0: 92}
                _ = v42
                _ = v43
            }
            v23 = "a"
        }
        if (v5 < (^(v5))) {
            var v44 string = v23
            _ = v44
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v3)
            for fi45 := 0; fi45 < 4; fi45++ {
                var v46 uint64 = (v3 + v3)
                var v47 float64 = ((((3.14159 * float64(1.5)) / -2.25) - float64(3.14159)) - 0.0)
                if v32 {
                    var v48 bool = true
                    if len(v4) > 0 {
                        v4[0] = ("\n\t" + "")
                    }
                    _ = v48
                }
                if (v2 > v2) {
                    fmt.Println(v5)
                }
                fmt.Println(v47, v46)
                _ = fi45
                _ = v46
                _ = v47
            }
        }
        _ = fi31
        _ = v32
    }
    var v49 int = 457
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v49
    _ = v5
}
