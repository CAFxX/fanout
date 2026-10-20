package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 uint; F2 bool }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 string; F1 bool }
func genFn0(p0 string, p1 uint64, p2 int64) uint {
    var v7 uint16 = 426
    fmt.Println(p1)
    var v8 int32 = -1
    switch int(v8) {
    case -153:
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := []int{}
            for fi11 := 0; fi11 < 19; fi11++ {
                switch p0 {
                case "foo bar":
                    _ = len(v10)
                    switch fi9 {
                    case 128:
                        v12 := GenS1{F0: 18446744073709551615, F1: 150}
                        _ = v12
                    case 0:
                        p0 = (p0 + "a")
                    }
                    v13 := GenS1{F0: 4294967295, F1: 259}
                    _ = v13
                default:
                    switch (fi9 | 579) {
                    case 956:
                        var v14 int64 = p2
                        _ = v14
                    case 127:
                        switch fi9 {
                        case 127:
                            v10 = append(v10, (fi11 | (fi9 >> ((18446744073709551615) & 63))))
                        default:
                            var v15 uint64 = (^(p1))
                            _ = v15
                        }
                    }
                }
                var v16 bool = false
                if ((127 != -111) && (p2 > 42)) {
                    var v17 bool = (65535 > 4294967295)
                    if len(v10) > 0 {
                        v10[0] = (fi9 << ((569) & 63))
                    }
                    var v18 int64 = (p2 | -9223372036854775808)
                    v10 = append(v10, fi11)
                    _ = v17
                    _ = v18
                }
                v19 := GenS2{F0: v8}
                _ = fi11
                _ = v16
                _ = v19
            }
            _ = fi9
            _ = v10
        }
    case 325:
        var v20 string = (p0 + p0)
        _ = v20
    case 1:
        v21 := GenS1{F0: 804, F1: 519}
        _ = v21
    default:
        var v22 uint8 = 0
        _ = v22
    }
    v7 += v7
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]uint32)
    var v3 uint64 = 2
    v4 := make(map[uint32]int)
    v5 := []string{("foo bar" + ("a" + "foo bar")), "a", ("\n\t" + "foo bar")}
    var v6 uint = 144
    _ = len(v4)
    v1 -= -387
    for fi23 := 0; fi23 < 6; fi23++ {
        for fi24 := 0; fi24 < 9; fi24++ {
            if len(v5) > 0 {
                v5[0] = ("\n\t" + "foo bar")
            }
            fmt.Println(fi23, fi24)
            fmt.Println(v1)
            _ = fi24
        }
        _ = fi23
    }
    v1 -= (-(v1))
    if len(v5) > 0 {
        v5[0] = ""
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        v6 += v6
        _ = fi25
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        if false {
            var v27 int8 = 1
            _ = v27
        }
        _ = fi26
    }
    v1 = (v1 ^ v1)
    if false {
        v5 = append(v5, "foo bar")
        for fi28 := 0; fi28 < 8; fi28++ {
            switch (^(fi28)) {
            case -2147483648:
                for fi29 := 0; fi29 < 1; fi29++ {
                    switch 1 {
                    case 128:
                        var v30 int = fi28
                        _ = v30
                    case 123:
                        v3 = v3
                    }
                    _ = fi29
                }
            case 611:
                fmt.Println(v1, v6, v6)
            case 2147483647:
                if false {
                    v6 -= 1
                    _ = len(v4)
                }
            }
            switch fi28 {
            case 32767:
                for fi31 := 0; fi31 < 3; fi31++ {
                    v1 -= v1
                    _ = fi31
                }
            case -445:
                v1 = 42
            case 2147483647:
                for fi32 := 0; fi32 < 10; fi32++ {
                    fmt.Println(v3, v3, v3)
                    _ = fi32
                }
            default:
                var v33 uint32 = uint32(v3)
                _ = v33
            }
            v3 -= (v3 ^ (v3 ^ (v3 | v3)))
            _ = fi28
        }
    } else {
        v6 = (v6 + 4294967295)
        if ("a" > "foo bar") {
            v1 = (v1 - 42)
            switch (("foo bar" + "") + "") {
            case "":
                v34 := GenS1{F0: v6, F1: 712}
                _ = v34
            case "foo bar":
                v1 = v1
            case "a":
                v35 := GenS3{F0: 926}
                _ = v35
            }
        } else {
            v36 := GenS4{F0: 42, F1: (v6 | (v6 - (^(v6)))), F2: (4294967295 <= 2147483647)}
            _ = v36
        }
    }
    v37 := GenS5{F0: 2, F1: ((1 == 42) && (2 <= 4294967295))}
    _ = len(v4)
    v38 := GenS6{F0: "a", F1: (-611 != 0)}
    v6 -= genFn0("", (v3 % ((4294967295) | 1)), v1)
    var v39 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
