package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16; F1 bool }
func genFn0(p0 string, p1 int, p2 int64) uint32 {
    for fi7 := 0; fi7 < 13; fi7++ {
        p0 = p0
        if false {
            p0 = "foo bar"
        }
        v8 := GenS1{F0: 0}
        _ = fi7
        _ = v8
    }
    var v9 string = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 42
}
func genFn1(p0 uint64) int {
    p0 += p0
    var v10 int32 = -1
    _ = p0
    _ = v10
    return 0
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{}
    var v3 uint32 = 0
    var v4 int32 = -888
    var v5 int64 = -306
    var v6 uint64 = 759
    if ((v3 | (v3 >> ((uint(v6)) & 63))) > (v3 / ((v3) | 1))) {
        v2 = append(v2, 128)
        var v11 int = 496
        delete(v1, 1)
        _ = v11
    }
    var v12 int = genFn1(1)
    var v13 uint16 = 65535
    if false {
        for fi14 := 0; fi14 < 10; fi14++ {
            v4 -= 0
            var v15 uint8 = 0
            var v16 int32 = 42
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    v17 := GenS1{F0: 1}
    v5 += (v5 / ((v5) | 1))
    var v18 float64 = ((0.0 - -2.25) / ((1.5 + (-2.25 + 3.14159)) * 0.1))
    if (!false) {
        v19 := GenS1{F0: v13}
        if (-128 == 127) {
            v12 = v12
        } else {
            var v20 float32 = float32(v18)
            if len(v2) > 0 {
                v2[0] = genFn1((v6 >> ((2) & 63)))
            }
            switch "a" {
            case "a":
                var v21 int64 = (v5 | v5)
                _ = v21
            case "\n\t":
                switch genFn1(v6) {
                case 214:
                    var v22 int8 = 0
                    _ = v22
                }
            case "":
                var v23 uint32 = (^(v3))
                _ = v23
            default:
                v13 = (v13 >> ((140) & 63))
            }
            _ = v20
        }
        v18 -= v18
        _ = v19
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 int64 = (v5 & v5)
        var v26 int16 = 32767
        v3 = v3
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        for fi28 := 0; fi28 < 7; fi28++ {
            v1[(v12 >> ((2) & 63))] = "foo bar"
            var v29 uint8 = 0
            v30 := GenS2{F0: ("hello" + "")}
            _ = fi28
            _ = v29
            _ = v30
        }
        var v31 int64 = int64(v12)
        _ = fi27
        _ = v31
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        var v33 uint16 = (v13 + (v13 & 32767))
        for fi34 := 0; fi34 < 19; fi34++ {
            if (("\n\t" + "hello") > "\n\t") {
                v18 = ((v18 * 1.5) * 1e10)
                v12 = fi32
            }
            v5 = (v5 ^ int64(v4))
            _ = fi34
        }
        _ = fi32
        _ = v33
    }
    _ = len(v2)
    for fi35 := 0; fi35 < 5; fi35++ {
        switch (^(v12)) {
        case 0:
            var v36 float64 = float64((0.1 - 1e10))
            _ = v36
        default:
            v1[-997] = (("" + "\n\t") + ((("foo bar" + "a") + ("" + "hello")) + "hello"))
        }
        v37 := GenS3{F0: genFn0("\n\t", v12, (v5 << ((255) & 63))), F1: 127, F2: 127}
        _ = fi35
        _ = v37
    }
    for fi38 := 0; fi38 < 14; fi38++ {
        var v39 uint64 = (^(v6))
        v40 := GenS4{F0: true}
        if (!(!false)) {
            for fi41 := 0; fi41 < 20; fi41++ {
                var v42 uint16 = v13
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if ((v4 / ((v4) | 1)) == (v4 * (-(v4)))) {
                    v1[int(v4)] = ""
                    v6 = v39
                    if (false && false) {
                        if ((!true) && ((((0.1 + 1.5) / 0.1) - 0.1) == -2.25)) {
                            if false {
                                var v43 int64 = (v5 & v5)
                                _ = v43
                            } else {
                                if (!(((586 >= 133) && false) || (!(-991 < (v4 % ((v4) | 1)))))) {
                                    v44 := GenS5{F0: (v39 | v6)}
                                    var v45 int = fi38
                                    var v46 bool = true
                                    v18 -= ((v18 * v18) + -2.25)
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                }
                            }
                        }
                    }
                }
                fmt.Println(v42, v39)
                var v47 int64 = (v5 + (^(v5)))
                _ = fi41
                _ = v42
                _ = v47
            }
        } else {
            v5 += (v5 & (v5 - 1))
            var v48 int16 = 1000
            v49 := GenS5{F0: (v39 & (v39 | (v6 & 982)))}
            _ = v48
            _ = v49
        }
        _ = fi38
        _ = v39
        _ = v40
    }
    if (!false) {
        var v50 bool = false
        switch "a" {
        case "foo bar":
            v51 := GenS5{F0: 426}
            _ = v51
        }
        _ = v50
    } else {
        for fi52 := 0; fi52 < 20; fi52++ {
            var v53 float64 = v18
            _ = fi52
            _ = v53
        }
    }
    if true {
        v4 += (v4 ^ v4)
        delete(v1, 4294967295)
    }
    v12 += (v12 >> ((255) & 63))
    v54 := GenS6{F0: 42, F1: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
}
