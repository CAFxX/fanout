package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint32; F1 bool; F2 int64 }
type GenS3 struct { F0 int; F1 int32 }
func genFn0(p0 uint32, p1 int16) int8 {
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := []int{}
        p1 = (p1 ^ 32767)
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi7
        _ = v8
    }
    if true {
        if (3.14159 < 1e10) {
            if (492 < 0) {
                var v9 string = ""
                _ = v9
            } else {
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(p1)
                }
            }
            v10 := []int{}
            _ = v10
        }
        var v11 uint = 895
        _ = v11
    } else {
        p1 = p1
        if false {
            switch -9223372036854775808 {
            case -9223372036854775808:
                p1 = (p1 + 42)
            case 218:
                if ((false && (((1 <= 42) && (42 <= 203)) && (42 > 87))) || true) {
                    var v12 bool = true
                    _ = v12
                } else {
                    v13 := []int{int(p0), 32768}
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    _ = v13
                }
            case 42:
                var v14 int = 127
                _ = v14
            }
            var v15 float32 = (0.0 * 1e10)
            fmt.Println(p0)
            _ = v15
        } else {
            for fi16 := 0; fi16 < 15; fi16++ {
                v17 := GenS1{F0: 68, F1: 0, F2: 2}
                v18 := GenS2{F0: p0, F1: false, F2: -959}
                _ = fi16
                _ = v17
                _ = v18
            }
            for fi19 := 0; fi19 < 6; fi19++ {
                if false {
                    v20 := []int{}
                    for fi21 := 0; fi21 < 13; fi21++ {
                        v22 := GenS2{F0: 42, F1: true, F2: 0}
                        p1 = p1
                        _ = fi21
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 20; fi23++ {
                        if false {
                            fmt.Println(p0, p1)
                            var v24 uint16 = 42
                            _ = v24
                        }
                        var v25 uint16 = 2
                        _ = fi23
                        _ = v25
                    }
                    _ = v20
                } else {
                    var v26 float32 = float32(0.1)
                    _ = v26
                }
                _ = fi19
            }
        }
        var v27 uint32 = (p0 + 914)
        _ = v27
    }
    v28 := GenS3{F0: -1, F1: 1}
    v29 := []int{}
    _ = p0
    _ = p1
    _ = v28
    _ = v29
    return 42
}

func main() {
    var v1 bool = false
    v2 := make(map[int]int)
    var v3 uint = 4294967295
    var v4 int8 = -63
    v5 := make(map[int]int)
    v6 := []string{"a"}
    var v30 int = 124
    for fi31 := 0; fi31 < 11; fi31++ {
        var v32 bool = v1
        if false {
            for fi33 := 0; fi33 < 20; fi33++ {
                v34 := GenS1{F0: 9223372036854775807, F1: 42, F2: 255}
                if (42 == -32768) {
                    if (742 > 455) {
                        v35 := GenS2{F0: 2, F1: (!true), F2: -9223372036854775808}
                        _ = v35
                    }
                    if len(v6) > 0 {
                        v6[0] = ""
                    }
                    v32 = false
                    fmt.Println(v1, v1)
                }
                _ = fi33
                _ = v34
            }
            v32 = v1
        } else {
            fmt.Println(fi31, v4)
            fmt.Println(v32, v3)
            fmt.Println(v1)
        }
        v3 += 42
        _ = fi31
        _ = v32
    }
    if len(v6) > 0 {
        v6[0] = ("\n\t" + "\n\t")
    }
    v1 = true
    var v36 float32 = 1e10
    for fi37 := 0; fi37 < 5; fi37++ {
        if v1 {
            v38 := GenS2{F0: 4294967295, F1: (170 > 1), F2: 639}
            if len(v6) > 0 {
                v6[0] = "a"
            }
            v4 += (v4 % (((^(v4))) | 1))
            _ = v38
        }
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v36, v36)
        case "":
            v39 := GenS1{F0: 1, F1: -843, F2: 0}
            _ = v39
        case "a":
            if (!(!(v3 > 105))) {
                var v40 int8 = -1
                fmt.Println(v40)
                _ = v40
            }
        }
        v4 = genFn0(uint32(v3), 42)
        _ = fi37
    }
    if v1 {
        if len(v6) > 0 {
            v6[0] = "foo bar"
        }
    } else {
        switch ("\n\t" + ("a" + ("foo bar" + "a"))) {
        case "hello":
            v41 := GenS2{F0: 1, F1: (2 < 98), F2: -334}
            _ = v41
        }
        v3 = 700
        fmt.Println(v3, v30)
    }
    fmt.Println(v30, v36)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
