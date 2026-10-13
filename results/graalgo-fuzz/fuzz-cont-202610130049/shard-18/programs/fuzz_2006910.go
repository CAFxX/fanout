package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int8 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint }
func genFn0() int8 {
    var v6 int = 0
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 int16 = -130
            switch "hello" {
            case "a":
                v10 := []int{fi7, (^(fi8)), fi8}
                _ = v10
            case "\n\t":
                v9 += v9
                fmt.Println(v6)
                v11 := []int{fi7, int(v9), fi7, fi7}
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    _ = v6
    return -1
}
func genFn1() uint64 {
    var v12 int16 = 0
    var v13 uint32 = 869
    var v14 bool = (true || false)
    for fi15 := 0; fi15 < 2; fi15++ {
        if v14 {
            if (fi15 < fi15) {
                v16 := GenS1{F0: ("hello" + "hello"), F1: -1, F2: 93}
                var v17 int = (^(fi15))
                _ = v16
                _ = v17
            } else {
                v12 = (v12 ^ (v12 - -32768))
                var v18 int64 = 9223372036854775807
                _ = v18
            }
        }
        _ = fi15
    }
    v14 = (2 <= 0)
    _ = v12
    _ = v13
    _ = v14
    return 42
}

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[uint32]int)
    var v3 bool = false
    var v4 uint16 = 2
    var v5 int8 = 42
    v19 := GenS2{F0: 127, F1: ""}
    switch "hello" {
    case "a":
        switch (("" + "") + "") {
        case "hello":
            for fi20 := 0; fi20 < 6; fi20++ {
                v3 = (!false)
                _ = fi20
            }
        case "\n\t":
            v21 := []int{-9223372036854775808, v1, (v1 % ((709) | 1))}
            _ = v21
        default:
            var v22 int64 = 463
            _ = v22
        }
    case "":
        fmt.Println(v1, v3, v4)
    default:
        var v23 uint32 = 1
        _ = v23
    }
    v3 = ((^(v4)) > (v4 - 65535))
    if ((!v3) && false) {
        var v24 int64 = 600
        if (v3 || v3) {
            var v25 uint32 = 42
            v25 -= (v25 << ((231) & 63))
            _ = v25
        } else {
            fmt.Println(v5, v1)
            v1 = (v1 % (((v1 % ((42) | 1))) | 1))
        }
        _ = v24
    } else {
        if ((855 == 157) && v3) {
            v5 += v5
            v26 := []int{v1, (v1 << ((65535) & 63))}
            var v27 uint = 869
            _ = v26
            _ = v27
        }
    }
    v5 = genFn0()
    var v28 int32 = -1
    v3 = ((v5 >= (v5 + genFn0())) || v3)
    var v29 string = "foo bar"
    var v30 float32 = -2.25
    if ("" < v29) {
        if (!(915 == 0)) {
            delete(v2, 2147483647)
        }
        if (true || v3) {
            if v3 {
                v31 := GenS2{F0: v1, F1: "hello"}
                v30 += v30
                v5 -= (v5 << ((4294967295) & 63))
                _ = v31
            } else {
                var v32 int = -1
                for fi33 := 0; fi33 < 5; fi33++ {
                    switch (v1 << ((18446744073709551615) & 63)) {
                    case 32767:
                        v34 := []int{fi33}
                        _ = v34
                    case -1:
                        v29 = v29
                    case 4294967295:
                        v1 -= v1
                    }
                    var v35 float64 = 0.0
                    v36 := []int{42, v32, 32767}
                    _ = fi33
                    _ = v35
                    _ = v36
                }
                v37 := GenS2{F0: v1, F1: "foo bar"}
                _ = v32
                _ = v37
            }
            if (v5 != (^(v5))) {
                v28 = (v28 << ((4294967295) & 63))
                if true {
                    v38 := GenS1{F0: v29, F1: -106, F2: (v5 << ((255) & 63))}
                    _ = v38
                } else {
                    v29 = (v29 + v29)
                }
                v5 -= 127
            } else {
                var v39 int8 = (v5 - (^(v5)))
                _ = v39
            }
            if false {
                var v40 int32 = (v28 << ((139) & 63))
                v1 -= -2147483648
                _ = v40
            }
        } else {
            v41 := GenS2{F0: -1, F1: v29}
            var v42 int64 = 316
            _ = v41
            _ = v42
        }
        fmt.Println(v4, v28)
    } else {
        for fi43 := 0; fi43 < 1; fi43++ {
            if (723 < 779) {
                for fi44 := 0; fi44 < 9; fi44++ {
                    var v45 uint = 4294967295
                    v46 := v2[2147483647]
                    var v47 uint32 = 2
                    _ = fi44
                    _ = v45
                    _ = v46
                    _ = v47
                }
            }
            _ = fi43
        }
        var v48 uint32 = 254
        _ = v48
    }
    v2[2147483647] = (v1 ^ 42)
    fmt.Println(v29, v28)
    fmt.Println(v3)
    if (860 == 2) {
        fmt.Println(v28, v3)
        var v49 int8 = 0
        _ = v49
    }
    switch (v29 + (v29 + (("\n\t" + (v29 + ("a" + v29))) + "hello"))) {
    case "a":
        if (!v3) {
            var v50 int = v1
            if true {
                var v51 float32 = ((v30 * float32(0.0)) * 1.5)
                _ = v51
            } else {
                if (!(!true)) {
                    v52 := GenS2{F0: (v50 | (-(v50))), F1: ""}
                    var v53 bool = v3
                    v29 = v29
                    _ = v52
                    _ = v53
                }
                var v54 uint = 18446744073709551615
                switch (v1 * (v1 - v50)) {
                case 0:
                    var v55 float64 = float64(v30)
                    _ = v55
                case 9223372036854775807:
                    v56 := GenS3{F0: (v54 + v54)}
                    _ = v56
                }
                _ = v54
            }
            _ = v50
        } else {
            v3 = true
            v4 = v4
            for fi57 := 0; fi57 < 4; fi57++ {
                var v58 uint16 = 0
                fmt.Println(fi57, v3)
                _ = fi57
                _ = v58
            }
        }
    default:
        delete(v2, 1)
    }
    v59 := []int{872, (^(v1))}
    for fi60 := 0; fi60 < 18; fi60++ {
        var v61 uint32 = 648
        for fi62 := 0; fi62 < 8; fi62++ {
            v63 := GenS1{F0: "", F1: -1, F2: (v5 % ((v5) | 1))}
            v30 += 0.0
            _ = fi62
            _ = v63
        }
        _ = fi60
        _ = v61
    }
    for fi64 := 0; fi64 < 20; fi64++ {
        for fi65 := 0; fi65 < 8; fi65++ {
            fmt.Println(v30, fi65)
            var v66 int64 = 42
            fmt.Println(fi64)
            _ = fi65
            _ = v66
        }
        v5 += genFn0()
        _ = fi64
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v59:", len(v59), v59)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v59
}
