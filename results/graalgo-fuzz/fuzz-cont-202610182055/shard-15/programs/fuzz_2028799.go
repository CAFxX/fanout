package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 string; F2 uint16 }
func genFn0(p0 bool, p1 int32, p2 uint) string {
    v7 := []int{42}
    var v8 int16 = 94
    switch "\n\t" {
    case "":
        switch (("" + "foo bar") + "a") {
        case "a":
            switch ("\n\t" + "") {
            case "":
                if true {
                    var v9 int64 = 0
                    if p0 {
                        p0 = false
                    } else {
                        v10 := GenS1{F0: ("hello" + "hello"), F1: 351}
                        _ = v10
                    }
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    _ = v9
                }
            case "a":
                v11 := GenS2{F0: 128}
                _ = v11
            case "foo bar":
                if (true || true) {
                    var v12 uint32 = 160
                    _ = v12
                } else {
                    if p0 {
                        var v13 bool = p0
                        var v14 int64 = 42
                        v14 = v14
                        _ = v13
                        _ = v14
                    }
                    switch 1 {
                    case 32768:
                        var v15 uint8 = 2
                        _ = v15
                    case 32767:
                        v16 := GenS3{F0: 20}
                        _ = v16
                    default:
                        if false {
                            _ = len(v7)
                            var v17 int8 = 66
                            var v18 float64 = ((1e10 * (1.5 / (-2.25 + 0.0))) - (-2.25 + -2.25))
                            _ = v17
                            _ = v18
                        } else {
                            switch -2147483648 {
                            case -2147483648:
                                var v19 bool = p0
                                _ = v19
                            }
                            p0 = p0
                        }
                    }
                }
            }
        default:
            p0 = false
        }
    case "hello":
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 uint32 = 2147483648
            p2 = p2
            _ = fi20
            _ = v21
        }
    case "a":
        switch (("" + "a") + "a") {
        case "foo bar":
            var v22 int16 = v8
            _ = v22
        case "":
            var v23 int = -325
            _ = v23
        case "a":
            fmt.Println(p0, p2)
        default:
            p2 -= (p2 + (^(p2)))
        }
    }
    p1 = 1
    for fi24 := 0; fi24 < 9; fi24++ {
        fmt.Println(fi24)
        _ = fi24
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ""
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]string)
    var v3 bool = (false || false)
    v4 := []uint32{1, 2147483648, 4294967295}
    var v5 float64 = (0.0 * (1e10 * -2.25))
    var v6 string = "\n\t"
    v4 = append(v4, 2147483648)
    for fi25 := 0; fi25 < 11; fi25++ {
        fmt.Println(fi25, v5)
        v26 := GenS4{F0: 42}
        var v27 uint16 = 32767
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        v29 := GenS3{F0: 100}
        _ = fi28
        _ = v29
    }
    v30 := GenS3{F0: 1}
    if (0 < 4294967295) {
        v31 := GenS2{F0: 127}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v32 uint = 4294967295
        _ = v31
        _ = v32
    }
    v6 = v6
    v3 = v3
    v6 = v6
    v2[174] = "foo bar"
    for fi33 := 0; fi33 < 15; fi33++ {
        v4 = append(v4, 148)
        var v34 uint8 = 2
        v3 = v3
        _ = fi33
        _ = v34
    }
    v35 := GenS3{F0: 0}
    var v36 uint8 = 255
    var v37 uint = 255
    v38 := GenS2{F0: 254}
    for fi39 := 0; fi39 < 17; fi39++ {
        fmt.Println(v5, v36)
        switch v6 {
        case "\n\t":
            v4 = append(v4, 42)
        case "":
            _ = len(v4)
        default:
            fmt.Println(v37)
        }
        if v3 {
            var v40 int16 = 0
            _ = v40
        } else {
            var v41 string = v6
            v3 = v3
            if true {
                v2[fi39] = "a"
                v42 := GenS5{F0: 2, F1: ((v41 + (v6 + v6)) + v6), F2: 2}
                if v3 {
                    v43 := v2[fi39]
                    if ((2 < 2147483647) && v3) {
                        var v44 string = ("foo bar" + (v41 + v43))
                        switch fi39 {
                        case -557:
                            v45 := GenS4{F0: 446}
                            _ = v45
                        case -9223372036854775808:
                            v36 = v36
                        }
                        _ = v44
                    } else {
                        _ = len(v1)
                        var v46 uint32 = 2
                        _ = v46
                    }
                    _ = v43
                }
                _ = v42
            }
            _ = v41
        }
        _ = fi39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
