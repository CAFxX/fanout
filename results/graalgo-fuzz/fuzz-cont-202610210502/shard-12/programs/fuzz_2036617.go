package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint16 }
type GenS2 struct { F0 int16; F1 int64; F2 int32 }
type GenS3 struct { F0 uint64 }
func genFn0() int32 {
    var v7 uint16 = 255
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := []int{}
        switch "a" {
        case "a":
            v10 := GenS1{F0: -876, F1: true, F2: (v7 / (((^(v7))) | 1))}
            _ = v10
        case "foo bar":
            v7 = (v7 >> ((1) & 63))
        case "hello":
            v7 = (v7 + (v7 % ((v7) | 1)))
        }
        v7 += v7
        _ = fi8
        _ = v9
    }
    if false {
        for fi11 := 0; fi11 < 2; fi11++ {
            fmt.Println(v7)
            v7 += v7
            _ = fi11
        }
    } else {
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS1{F0: 42, F1: false, F2: v7}
            var v14 uint = 1
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    _ = v7
    return -1
}
func genFn1(p0 uint16, p1 uint8, p2 string) bool {
    p1 += 255
    for fi15 := 0; fi15 < 16; fi15++ {
        for fi16 := 0; fi16 < 16; fi16++ {
            v17 := []int{-592, fi15, fi16, 827}
            p0 -= 426
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v18 := GenS1{F0: 1, F1: true, F2: p0}
    if true {
        v19 := []int{}
        _ = v19
    }
    if true {
        for fi20 := 0; fi20 < 5; fi20++ {
            for fi21 := 0; fi21 < 3; fi21++ {
                if false {
                    v22 := GenS1{F0: 0, F1: (p0 > 345), F2: (^(p0))}
                    p0 -= p0
                    _ = v22
                } else {
                    var v23 int8 = 79
                    fmt.Println(fi20)
                    _ = v23
                }
                var v24 int64 = -948
                v25 := GenS2{F0: 0, F1: (^(v24)), F2: genFn0()}
                _ = fi21
                _ = v24
                _ = v25
            }
            _ = fi20
        }
        var v26 uint32 = 2147483648
        if false {
            p0 += (p0 - 760)
        }
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return (!false)
}

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]int)
    v3 := []string{("" + ("" + "a")), "foo bar"}
    v4 := make(map[string]string)
    var v5 string = "foo bar"
    var v6 int64 = 0
    if false {
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 bool = false
            _ = len(v2)
            for fi29 := 0; fi29 < 12; fi29++ {
                var v30 int8 = 28
                _ = fi29
                _ = v30
            }
            _ = fi27
            _ = v28
        }
        if (65535 >= 17) {
            v31 := GenS1{F0: 297, F1: (!true), F2: 2}
            v32 := GenS1{F0: (v6 << ((0) & 63)), F1: ((1 <= 42) || true), F2: 32767}
            _ = v31
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        var v34 uint = v1
        if (fi33 <= (fi33 ^ (fi33 >> ((v34) & 63)))) {
            _ = len(v3)
            for fi35 := 0; fi35 < 13; fi35++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v34 += v1
                _ = fi35
            }
            var v36 int16 = 1000
            _ = v36
        } else {
            v37 := GenS3{F0: 677}
            if false {
                var v38 int64 = (^(v6))
                v34 += (v1 % (((v34 + (v1 * v34))) | 1))
                v1 = v34
                _ = v38
            } else {
                var v39 int = (fi33 >> ((v34) & 63))
                v6 = (v6 | v6)
                switch ((v5 + "a") + ("\n\t" + "hello")) {
                case "\n\t":
                    v40 := GenS2{F0: -1, F1: v6, F2: int32(v6)}
                    _ = v40
                case "":
                    v5 = v5
                }
                _ = v39
            }
            _ = v37
        }
        switch v5 {
        case "foo bar":
            v4[""] = "a"
        }
        _ = fi33
        _ = v34
    }
    var v41 int64 = (-(v6))
    v3 = append(v3, ("a" + v5))
    v42 := GenS3{F0: 4294967295}
    var v43 int32 = genFn0()
    var v44 float32 = 0.1
    var v45 uint32 = 0
    switch 548 {
    case 32767:
        _ = len(v4)
    case 2147483647:
        v46 := GenS3{F0: 18446744073709551615}
        _ = v46
    case -967:
        var v47 int8 = 109
        _ = v47
    }
    var v48 float32 = v44
    v2[v45] = 0
    if (-128 != 0) {
        fmt.Println(v44, v41)
        var v49 int16 = int16(v43)
        _ = v49
    }
    if len(v3) > 0 {
        v3[0] = ""
    }
    fmt.Println(v41, v44)
    if len(v3) > 0 {
        v3[0] = ("hello" + v5)
    }
    var v50 int8 = 127
    v51 := GenS3{F0: 4294967295}
    v1 -= 253
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v48
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
