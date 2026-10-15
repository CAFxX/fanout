package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint32, p1 int16) uint8 {
    v6 := GenS1{F0: -128}
    p0 += p0
    fmt.Println(p0, p0)
    if false {
        var v7 bool = false
        _ = v7
    }
    p1 += (p1 ^ 1)
    _ = p0
    _ = p1
    _ = v6
    return 255
}
func genFn1() int {
    for fi8 := 0; fi8 < 11; fi8++ {
        switch ("hello" + "\n\t") {
        case "":
            v9 := []int{-1, (fi8 & (-(fi8)))}
            _ = v9
        case "hello":
            fmt.Println(fi8)
        case "a":
            var v10 int8 = 42
            _ = v10
        default:
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 string = ("a" + (("\n\t" + ("hello" + "hello")) + "hello"))
                v13 := []int{fi11, -2147483648, fi11}
                var v14 uint16 = 693
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
        }
        v15 := []int{4294967295, fi8}
        _ = fi8
        _ = v15
    }
    var v16 bool = (-99 >= 1)
    if true {
        var v17 int = -9223372036854775808
        var v18 int8 = -128
        v19 := []int{v17}
        _ = v17
        _ = v18
        _ = v19
    }
    _ = v16
    return 32768
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[int]int)
    var v3 bool = (v1 != v1)
    v4 := make(map[int]string)
    v5 := []string{"a"}
    v20 := GenS2{F0: 856}
    fmt.Println(v1)
    fmt.Println(v3, v1)
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            v23 := GenS1{F0: 114}
            var v24 int64 = -157
            if (!v3) {
                if true {
                    var v25 uint64 = v1
                    if v3 {
                        var v26 uint16 = 0
                        _ = v26
                    }
                    _ = v25
                } else {
                    v27 := GenS3{F0: fi21}
                    v24 -= v24
                    fmt.Println(fi21)
                    switch 1 {
                    case 42:
                        var v28 float64 = 1e10
                        _ = v28
                    case 32767:
                        var v29 float32 = float32((0.0 * 1.5))
                        _ = v29
                    default:
                        switch genFn1() {
                        case 127:
                            v30 := GenS4{F0: 0, F1: (-35 == 21)}
                            _ = v30
                        default:
                            fmt.Println(fi21, v3, fi22)
                        }
                    }
                    v3 = v3
                    v3 = v3
                    _ = v27
                }
                v3 = true
                switch fi21 {
                case 909:
                    switch "\n\t" {
                    case "foo bar":
                        v3 = (!(int8(v1) >= -128))
                    }
                    v5 = append(v5, "a")
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                switch ("a" + "\n\t") {
                case "\n\t":
                    v31 := GenS2{F0: 728}
                    _ = v31
                case "a":
                    fmt.Println(fi22, fi21)
                case "":
                    v5 = append(v5, ("foo bar" + ("a" + ("foo bar" + "\n\t"))))
                }
            } else {
                if v3 {
                    fmt.Println(v24, fi22, fi21)
                } else {
                    v2[fi22] = (fi21 % ((fi21) | 1))
                }
                var v32 string = "a"
                _ = v32
            }
            v33 := GenS4{F0: 32767, F1: v3}
            _ = len(v5)
            v34 := GenS3{F0: (fi21 << ((2) & 63))}
            _ = fi22
            _ = v23
            _ = v24
            _ = v33
            _ = v34
        }
        v35 := GenS2{F0: 32767}
        switch (^(fi21)) {
        case -978:
            v36 := GenS2{F0: 32767}
            _ = v36
        case 852:
            fmt.Println(v3)
        default:
            v3 = v3
        }
        _ = fi21
        _ = v35
    }
    var v37 string = (("foo bar" + "a") + "foo bar")
    _ = len(v5)
    if (v3 && (-2.25 != 1.5)) {
        switch -2147483648 {
        case -1:
            v37 = v37
        case 4294967295:
            fmt.Println(v37)
        case 219:
            var v38 float32 = 0.0
            _ = v38
        }
        v39 := GenS5{F0: -11}
        for fi40 := 0; fi40 < 16; fi40++ {
            if len(v5) > 0 {
                v5[0] = v37
            }
            delete(v4, (^(fi40)))
            _ = fi40
        }
        _ = v39
    } else {
        var v41 float32 = 1e10
        if (false && (!((227 <= 42) || (false && (!(int64(v1) >= 1)))))) {
            v3 = false
            if false {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                fmt.Println(v41, v41, v37)
                v42 := v2[32767]
                _ = v42
            }
        }
        _ = v41
    }
    delete(v4, -9223372036854775808)
    if true {
        fmt.Println(v1)
    } else {
        if false {
            fmt.Println(v37)
            v1 -= 994
            v37 = ("\n\t" + (("a" + v37) + "a"))
        } else {
            if (!v3) {
                for fi43 := 0; fi43 < 19; fi43++ {
                    if v3 {
                        if v3 {
                            var v44 uint16 = 864
                            if false {
                                var v45 string = (v37 + v37)
                                v37 = v37
                                _ = v45
                            }
                            _ = v44
                        }
                        v1 += 2
                        v4[(-(fi43))] = "a"
                    }
                    v46 := v4[0]
                    _ = fi43
                    _ = v46
                }
                v47 := GenS6{F0: 796}
                v48 := GenS5{F0: -128}
                _ = v47
                _ = v48
            } else {
                fmt.Println(v3)
                var v49 uint = 18446744073709551615
                fmt.Println(v1, v3)
                _ = v49
            }
            var v50 uint8 = 255
            v51 := GenS1{F0: 1}
            _ = v50
            _ = v51
        }
        v5 = append(v5, ((v37 + "foo bar") + "hello"))
        if v3 {
            fmt.Println(v37, v1)
            fmt.Println(v37, v3, v37)
        } else {
            switch v37 {
            case "a":
                var v52 int8 = 0
                _ = v52
            case "\n\t":
                v3 = v3
            }
            v3 = ("a" >= v37)
            v53 := GenS5{F0: 1}
            _ = v53
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
