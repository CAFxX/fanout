package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 int32; F1 int32; F2 int64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS9 struct { F0 int64; F1 int16 }
type GenS10 struct { F0 int32 }
func genFn0(p0 uint16, p1 uint16, p2 uint32) uint {
    p1 = p0
    for fi4 := 0; fi4 < 10; fi4++ {
        v5 := GenS1{F0: 18446744073709551615, F1: ("foo bar" + "a"), F2: 42}
        _ = fi4
        _ = v5
    }
    var v6 uint = 2
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 int64, p1 uint64) string {
    if (507 != 0) {
        v7 := []int{32768}
        p0 -= (p0 / ((p0) | 1))
        _ = v7
    } else {
        v8 := []int{-212, 32768}
        v9 := GenS2{F0: (("foo bar" + "") + "a")}
        _ = v8
        _ = v9
    }
    fmt.Println(p0, p1)
    p0 = -9223372036854775808
    v10 := []int{32767}
    _ = p0
    _ = p1
    _ = v10
    return (("\n\t" + ("\n\t" + "hello")) + "hello")
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 2
    v3 := []uint32{65535}
    var v11 int32 = -1
    if (!(127 > 0)) {
        if true {
            var v12 uint = 4294967295
            _ = v12
        }
        var v13 uint32 = v2
        v14 := GenS1{F0: 490, F1: "foo bar", F2: 18446744073709551615}
        _ = v13
        _ = v14
    }
    v2 += 640
    for fi15 := 0; fi15 < 12; fi15++ {
        v2 = (^(v2))
        _ = fi15
    }
    v2 = 2147483647
    for fi16 := 0; fi16 < 19; fi16++ {
        if len(v3) > 0 {
            v3[0] = (^(v2))
        }
        _ = fi16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        v1[(fi17 - fi17)] = fi17
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 int64 = -1
            if false {
                v19 -= (v19 + (v19 & -634))
                v20 := GenS1{F0: 18446744073709551615, F1: ("a" + ""), F2: 4294967295}
                _ = v20
            } else {
                v11 = -1
                switch ("a" + (("a" + ((("hello" + "") + "a") + "hello")) + "")) {
                case "foo bar":
                    fmt.Println(fi18, v19, v11)
                case "hello":
                    var v21 float32 = (1e10 + float32(1e10))
                    _ = v21
                }
            }
            v22 := GenS3{F0: 0}
            v19 = v19
            fmt.Println(fi17, fi17)
            fmt.Println(v2, v11, v11)
            v23 := GenS4{F0: 234}
            v1[(fi17 & 0)] = (fi17 / ((fi18) | 1))
            var v24 float64 = 1.5
            _ = fi18
            _ = v19
            _ = v22
            _ = v23
            _ = v24
        }
        _ = fi17
    }
    var v25 uint64 = 963
    v26 := GenS5{F0: v11, F1: 18446744073709551615}
    switch "\n\t" {
    case "foo bar":
        for fi27 := 0; fi27 < 15; fi27++ {
            fmt.Println(v2, v2, v2)
            _ = fi27
        }
    case "\n\t":
        fmt.Println(v25, v2)
    default:
        for fi28 := 0; fi28 < 6; fi28++ {
            if (fi28 != (fi28 / ((fi28) | 1))) {
                var v29 uint = uint(v2)
                _ = v29
            } else {
                fmt.Println(v11, v2)
                if true {
                    fmt.Println(v25, v2)
                    v30 := GenS6{F0: (v11 % ((2147483647) | 1)), F1: v11, F2: 9223372036854775807}
                    _ = v30
                }
                v1[fi28] = fi28
            }
            if false {
                v11 = int32(v2)
                for fi31 := 0; fi31 < 15; fi31++ {
                    v32 := GenS1{F0: 1, F1: "a", F2: 4294967295}
                    _ = fi31
                    _ = v32
                }
            }
            if false {
                for fi33 := 0; fi33 < 19; fi33++ {
                    _ = len(v3)
                    v34 := GenS7{F0: (^(v25))}
                    _ = fi33
                    _ = v34
                }
                if true {
                    var v35 int = 195
                    _ = len(v3)
                    _ = v35
                }
                var v36 bool = true
                _ = v36
            }
            _ = fi28
        }
    }
    var v37 uint8 = 42
    v11 += v11
    v38 := GenS7{F0: (v25 / (((^(v25))) | 1))}
    for fi39 := 0; fi39 < 10; fi39++ {
        switch "hello" {
        case "":
            v40 := v1[fi39]
            _ = v40
        case "foo bar":
            v2 -= v2
        case "a":
            switch ((("hello" + "hello") + "") + (("hello" + ("a" + ("a" + "\n\t"))) + (("hello" + "a") + "hello"))) {
            case "foo bar":
                fmt.Println(fi39, v37, v2)
            }
        default:
            var v41 bool = (((float64(1.5) == 0.0) && (435 > 255)) && (0 > 141))
            _ = v41
        }
        switch "\n\t" {
        case "\n\t":
            v42 := GenS8{F0: (v2 % ((v2) | 1)), F1: 42, F2: 0}
            _ = v42
        }
        if (-32768 <= 42) {
            if true {
                for fi43 := 0; fi43 < 13; fi43++ {
                    v44 := GenS7{F0: v25}
                    fmt.Println(v25)
                    _ = fi43
                    _ = v44
                }
            }
        } else {
            v45 := GenS9{F0: -673, F1: 32767}
            fmt.Println(v37)
            v46 := GenS10{F0: (-(v11))}
            _ = v45
            _ = v46
        }
        _ = fi39
    }
    for fi47 := 0; fi47 < 15; fi47++ {
        fmt.Println(v11, v11, fi47)
        for fi48 := 0; fi48 < 20; fi48++ {
            var v49 int = fi47
            switch ("foo bar" + "") {
            case "foo bar":
                var v50 uint = 18446744073709551615
                _ = v50
            case "\n\t":
                v1[1] = (v49 - (v49 / ((-9223372036854775808) | 1)))
            }
            _ = fi48
            _ = v49
        }
        v37 -= v37
        _ = fi47
    }
    if (((-(v11)) != v11) && (1 <= (v25 + 725))) {
        fmt.Println(v11)
    } else {
        v3 = append(v3, v2)
        if (v37 == uint8(v2)) {
            var v51 int16 = -1
            _ = v51
        }
    }
    fmt.Println(v11, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v38
}
