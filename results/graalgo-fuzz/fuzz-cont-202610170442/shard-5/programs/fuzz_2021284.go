package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int32; F2 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 bool; F1 bool }
func genFn0(p0 string, p1 int8) int8 {
    for fi4 := 0; fi4 < 17; fi4++ {
        p0 = ""
        fmt.Println(p0, p1, p1)
        switch fi4 {
        case -373:
            if true {
                p0 = (("" + "hello") + p0)
            }
        }
        _ = fi4
    }
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := []int{(fi5 ^ fi5), 0, (fi5 & (fi5 ^ fi5))}
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: -87, F1: 32767}
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int = 128
            _ = fi9
            _ = v10
        }
        fmt.Println(p1)
        var v11 int64 = 42
        _ = fi8
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    return 127
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 1
    v3 := make(map[uint32]string)
    v12 := []int{-598, -91, 42, 0}
    v13 := GenS2{F0: 0}
    var v14 uint16 = 32767
    var v15 bool = ((("hello" + "\n\t") <= ("\n\t" + "hello")) && true)
    _ = len(v1)
    if v15 {
        v2 = (v2 / ((v2) | 1))
        v16 := GenS2{F0: -1}
        _ = v16
    }
    switch "" {
    case "foo bar":
        v2 -= v2
    case "":
        if (v15 && (1.5 < ((0.1 / (0.0 + 3.14159)) + 1.5))) {
            v14 = (v14 % ((v14) | 1))
            var v17 int16 = 32767
            var v18 bool = v15
            _ = v17
            _ = v18
        }
    case "a":
        v14 -= v14
    }
    v2 -= 42
    for fi19 := 0; fi19 < 11; fi19++ {
        fmt.Println(v2)
        var v20 uint16 = (v14 / ((242) | 1))
        _ = fi19
        _ = v20
    }
    v15 = (79 <= 42)
    if v15 {
        fmt.Println(v14, v14, v2)
        for fi21 := 0; fi21 < 1; fi21++ {
            if (true && (-1 != 42)) {
                _ = len(v3)
                v14 = v14
                v15 = (216 >= 42)
            } else {
                for fi22 := 0; fi22 < 2; fi22++ {
                    v23 := GenS1{F0: 1, F1: fi22}
                    fmt.Println(fi21)
                    switch (fi21 >> ((4294967295) & 63)) {
                    case 0:
                        v2 = (^(v2))
                    default:
                        if v15 {
                            v12 = append(v12, fi21)
                            v24 := GenS2{F0: 32767}
                            if v15 {
                                v25 := GenS3{F0: (v14 % ((v14) | 1)), F1: -1, F2: "\n\t"}
                                v26 := GenS4{F0: (v2 % ((v2) | 1))}
                                if v15 {
                                    if false {
                                        _ = len(v12)
                                        var v27 float32 = ((1e10 - -2.25) + (((0.1 / float32(0.1)) / float32(1.5)) / 0.1))
                                        _ = v27
                                    }
                                    var v28 string = ("foo bar" + ("a" + (("foo bar" + "a") + "hello")))
                                    _ = v28
                                } else {
                                    fmt.Println(v2)
                                    fmt.Println(fi22, v14)
                                    v12 = append(v12, (fi22 % (((fi21 | (fi21 * (fi21 ^ (^(fi22)))))) | 1)))
                                    v15 = (247 == 128)
                                    _ = len(v12)
                                }
                                _ = v25
                                _ = v26
                            } else {
                                switch "hello" {
                                case "a":
                                    switch (fi22 % ((947) | 1)) {
                                    case -906:
                                        v29 := GenS5{F0: 638, F1: 2147483648}
                                        _ = v29
                                    case -9223372036854775808:
                                        switch "hello" {
                                        case "hello":
                                            var v30 int = fi21
                                            _ = v30
                                        }
                                    }
                                case "\n\t":
                                    delete(v1, ("foo bar" + ("hello" + "hello")))
                                case "foo bar":
                                    var v31 uint = 18446744073709551615
                                    _ = v31
                                }
                                var v32 int32 = 985
                                _ = len(v12)
                                _ = v32
                            }
                            v2 = (^(v2))
                            if true {
                                var v33 string = ""
                                var v34 float32 = (1e10 - 3.14159)
                                if v15 {
                                    if v15 {
                                        v15 = (v15 && true)
                                        v33 = "hello"
                                        v35 := v3[2147483647]
                                        _ = v35
                                    }
                                }
                                _ = v33
                                _ = v34
                            }
                            _ = v24
                        }
                    }
                    _ = fi22
                    _ = v23
                }
            }
            v36 := GenS6{F0: (-9223372036854775808 < -324), F1: true}
            var v37 int16 = 0
            _ = fi21
            _ = v36
            _ = v37
        }
        for fi38 := 0; fi38 < 5; fi38++ {
            var v39 float32 = (-2.25 * 1.5)
            var v40 bool = (((^(v2)) == (v2 + (v2 & (^(v2))))) && false)
            _ = fi38
            _ = v39
            _ = v40
        }
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    switch "hello" {
    case "":
        var v41 uint = 728
        _ = v41
    default:
        switch "a" {
        case "a":
            fmt.Println(v15)
        case "":
            var v42 int8 = 74
            _ = v42
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
