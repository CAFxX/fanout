package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int32; F1 int64; F2 uint64 }
type GenS4 struct { F0 int8; F1 uint; F2 bool }
func genFn0() int64 {
    var v4 uint8 = 2
    if ((-1 >= 948) && (float32(1.5) == (3.14159 - 0.0))) {
        v4 = (v4 | v4)
        for fi5 := 0; fi5 < 17; fi5++ {
            v6 := []int{fi5, fi5}
            for fi7 := 0; fi7 < 15; fi7++ {
                var v8 int64 = 42
                if ((true || true) && (1 != 18446744073709551615)) {
                    _ = len(v6)
                    var v9 uint16 = 722
                    _ = v9
                } else {
                    v6 = append(v6, (fi5 << ((42) & 63)))
                    var v10 int8 = -68
                    _ = v10
                }
                if len(v6) > 0 {
                    v6[0] = fi7
                }
                _ = fi7
                _ = v8
            }
            var v11 string = "foo bar"
            _ = fi5
            _ = v11
            _ = v6
        }
    }
    _ = v4
    return -9223372036854775808
}
func genFn1(p0 string, p1 uint, p2 int32) uint64 {
    if (-543 < -9223372036854775808) {
        switch "foo bar" {
        case "":
            switch -9223372036854775808 {
            case 32768:
                p0 = (p0 + "foo bar")
            case 627:
                for fi12 := 0; fi12 < 13; fi12++ {
                    var v13 string = ((p0 + p0) + p0)
                    _ = fi12
                    _ = v13
                }
            }
        case "foo bar":
            fmt.Println(p2)
        case "a":
            var v14 uint = p1
            _ = v14
        }
        if true {
            v15 := GenS1{F0: false, F1: 100}
            fmt.Println(p2, p2, p1)
            _ = v15
        }
    }
    p0 = (p0 + p0)
    for fi16 := 0; fi16 < 18; fi16++ {
        if true {
            for fi17 := 0; fi17 < 15; fi17++ {
                p0 = "hello"
                var v18 uint8 = 46
                switch -9223372036854775808 {
                case 512:
                    v19 := GenS1{F0: ("" != ((p0 + (p0 + p0)) + p0)), F1: -128}
                    _ = v19
                case 4294967295:
                    var v20 int64 = 1
                    _ = v20
                }
                _ = fi17
                _ = v18
            }
            v21 := GenS2{F0: -9223372036854775808, F1: (^(fi16))}
            _ = v21
        } else {
            v22 := GenS2{F0: -807, F1: -379}
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := GenS1{F0: ((p0 + (((p0 + p0) + p0) + ("" + p0))) < ("hello" + p0)), F1: 42}
                _ = fi23
                _ = v24
            }
            _ = v22
        }
        var v25 int8 = 127
        _ = fi16
        _ = v25
    }
    p0 = "foo bar"
    var v26 uint64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return v26
}

func main() {
    var v1 bool = false
    v2 := []string{("hello" + ("\n\t" + "a"))}
    v3 := make(map[int]int)
    var v27 uint32 = 655
    v1 = (1.5 > 1e10)
    if v1 {
        for fi28 := 0; fi28 < 15; fi28++ {
            var v29 bool = (true || true)
            v30 := GenS3{F0: -974, F1: -284, F2: 4294967295}
            for fi31 := 0; fi31 < 8; fi31++ {
                var v32 bool = (2 < 18446744073709551615)
                _ = fi31
                _ = v32
            }
            _ = fi28
            _ = v29
            _ = v30
        }
    } else {
        for fi33 := 0; fi33 < 18; fi33++ {
            v2 = append(v2, ("foo bar" + ("a" + (("" + ("" + "\n\t")) + ("\n\t" + "")))))
            _ = fi33
        }
        if false {
            if len(v2) > 0 {
                v2[0] = ""
            }
            if ("\n\t" == ("hello" + ("a" + "a"))) {
                v34 := v3[0]
                v34 -= 32767
                v35 := GenS4{F0: 127, F1: 567, F2: (32767 < 1)}
                _ = v34
                _ = v35
            } else {
                v3[0] = 32767
                if v1 {
                    v36 := v3[1]
                    _ = v36
                } else {
                    delete(v3, 148)
                    v2 = append(v2, "hello")
                }
            }
            v37 := GenS4{F0: 0, F1: 1, F2: (("\n\t" + "\n\t") < "a")}
            _ = v37
        }
        var v38 uint64 = 4294967295
        _ = v38
    }
    fmt.Println(v27)
    v27 -= v27
    v1 = true
    var v39 int16 = 515
    for fi40 := 0; fi40 < 6; fi40++ {
        var v41 bool = v1
        _ = fi40
        _ = v41
    }
    if v1 {
        fmt.Println(v1, v1, v27)
        var v42 uint64 = 1
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v39
}
