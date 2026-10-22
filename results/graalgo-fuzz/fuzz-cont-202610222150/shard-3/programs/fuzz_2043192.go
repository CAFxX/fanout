package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 uint }
func genFn0(p0 uint64, p1 uint8, p2 string) int32 {
    var v7 int64 = 1
    switch 271 {
    case -1:
        v8 := []int{127, 32768}
        _ = v8
    case 4294967295:
        var v9 float64 = 0.0
        _ = v9
    case 1:
        fmt.Println(p0)
    default:
        switch (p2 + p2) {
        case "a":
            p1 = p1
        case "\n\t":
            v10 := GenS1{F0: p1}
            _ = v10
        }
    }
    v11 := GenS2{F0: p2}
    if (!true) {
        var v12 int32 = 0
        if true {
            v7 = (v7 - (^(v7)))
            for fi13 := 0; fi13 < 17; fi13++ {
                v14 := GenS2{F0: ("hello" + p2)}
                v15 := []int{}
                _ = fi13
                _ = v14
                _ = v15
            }
        } else {
            var v16 uint = 18446744073709551615
            p2 = "foo bar"
            _ = v16
        }
        if false {
            v7 = v7
            p2 = p2
        }
        _ = v12
    } else {
        v17 := GenS2{F0: p2}
        for fi18 := 0; fi18 < 2; fi18++ {
            p1 += p1
            _ = fi18
        }
        for fi19 := 0; fi19 < 15; fi19++ {
            if true {
                var v20 int = 0
                for fi21 := 0; fi21 < 20; fi21++ {
                    p0 = p0
                    _ = fi21
                }
                p2 = p2
                _ = v20
            } else {
                var v22 uint = 42
                _ = v22
            }
            _ = fi19
        }
        _ = v17
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        fmt.Println(fi23, v7, fi23)
        v24 := []int{}
        var v25 int8 = 127
        _ = fi23
        _ = v24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return 42
}

func main() {
    var v1 bool = (!true)
    var v2 uint32 = 614
    var v3 int = -193
    v4 := make(map[int]bool)
    v5 := make(map[int]uint32)
    v6 := []string{"", (("\n\t" + "a") + "hello"), ("foo bar" + "")}
    var v26 uint = 18446744073709551615
    switch "\n\t" {
    case "\n\t":
        v27 := v4[v3]
        _ = v27
    case "a":
        for fi28 := 0; fi28 < 5; fi28++ {
            switch (("a" + ("" + "\n\t")) + ("" + ("\n\t" + "foo bar"))) {
            case "foo bar":
                v29 := GenS3{F0: false}
                _ = v29
            case "\n\t":
                switch (v3 * (v3 ^ 962)) {
                case -686:
                    v2 = 2147483647
                case -9223372036854775808:
                    if ((float32((0.0 - (((0.1 / -2.25) + 0.1) + (3.14159 * (0.0 * 1e10))))) * 0.0) <= (1.5 + -2.25)) {
                        v30 := GenS3{F0: (-128 < 127)}
                        switch v3 {
                        case -9223372036854775808:
                            if (!v1) {
                                v26 = (^(v26))
                            }
                        case -1:
                            if true {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                for fi31 := 0; fi31 < 2; fi31++ {
                                    if v1 {
                                        v32 := GenS2{F0: "\n\t"}
                                        var v33 uint64 = 18446744073709551615
                                        v34 := v5[(v3 | (fi28 / ((v3) | 1)))]
                                        if len(v6) > 0 {
                                            v6[0] = "foo bar"
                                        }
                                        if true {
                                            var v35 uint64 = v33
                                            _ = v35
                                        } else {
                                            v6 = append(v6, "\n\t")
                                        }
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    }
                                    v3 += (fi31 & v3)
                                    _ = fi31
                                }
                                var v36 uint8 = 193
                                _ = v36
                            } else {
                                var v37 float32 = float32(((float64((0.0 / -2.25)) * 3.14159) / (1.5 * -2.25)))
                                _ = v37
                            }
                        case -2147483648:
                            var v38 float64 = 1e10
                            _ = v38
                        }
                        for fi39 := 0; fi39 < 1; fi39++ {
                            var v40 uint8 = 154
                            v41 := GenS1{F0: v40}
                            v42 := GenS2{F0: "a"}
                            if false {
                                v43 := GenS4{F0: "foo bar", F1: (v40 - 219)}
                                _ = v43
                            }
                            v40 = (v40 & (^(v40)))
                            _ = fi39
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                        _ = v30
                    }
                default:
                    for fi44 := 0; fi44 < 9; fi44++ {
                        fmt.Println(v26, v1)
                        v45 := GenS5{F0: 0, F1: 127}
                        v46 := v4[(v3 & 128)]
                        v47 := GenS2{F0: (("\n\t" + "\n\t") + "\n\t")}
                        _ = fi44
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                }
            case "a":
                v48 := GenS6{F0: 375}
                _ = v48
            }
            switch ("\n\t" + "a") {
            case "a":
                fmt.Println(v3, fi28)
            case "foo bar":
                v26 = 255
            default:
                var v49 int64 = 1
                _ = v49
            }
            for fi50 := 0; fi50 < 18; fi50++ {
                fmt.Println(fi28, v2)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = len(v6)
                if true {
                    var v51 uint32 = 42
                    fmt.Println(v51, v1, v51)
                    _ = v51
                }
                var v52 uint32 = (^(v2))
                _ = fi50
                _ = v52
            }
            _ = fi28
        }
    default:
        v3 = v3
    }
    v26 -= v26
    delete(v4, (v3 + (v3 & (v3 / ((128) | 1)))))
    var v53 int8 = 100
    var v54 uint = (v26 << ((710) & 63))
    if ((-523 == 1) && v1) {
        var v55 int8 = v53
        var v56 uint = 0
        _ = v55
        _ = v56
    } else {
        var v57 bool = v1
        var v58 uint8 = 42
        _ = v57
        _ = v58
    }
    if true {
        for fi59 := 0; fi59 < 16; fi59++ {
            if v1 {
                for fi60 := 0; fi60 < 7; fi60++ {
                    var v61 int = 0
                    fmt.Println(v54, v1, v53)
                    _ = fi60
                    _ = v61
                }
                v6 = append(v6, ("hello" + "hello"))
            }
            fmt.Println(v54)
            v54 = 0
            _ = fi59
        }
        if true {
            var v62 uint = v26
            for fi63 := 0; fi63 < 19; fi63++ {
                for fi64 := 0; fi64 < 4; fi64++ {
                    fmt.Println(v62, v54, fi63)
                    v54 += (v26 ^ 255)
                    fmt.Println(v3, v2, v3)
                    _ = fi64
                }
                if (1 == 42) {
                    var v65 uint32 = v2
                    _ = v65
                }
                _ = fi63
            }
            _ = v62
        }
    } else {
        var v66 string = "hello"
        _ = v66
    }
    delete(v5, int(v54))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
