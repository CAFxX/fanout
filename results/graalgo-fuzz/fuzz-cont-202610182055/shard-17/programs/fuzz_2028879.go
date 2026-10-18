package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int; F2 int64 }
type GenS4 struct { F0 string; F1 uint8 }

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]bool)
    v3 := make(map[uint32]string)
    var v4 float32 = (0.0 + -2.25)
    var v5 int64 = 1
    var v6 uint64 = 1
    var v7 uint32 = 42
    fmt.Println(v6, v5)
    for fi8 := 0; fi8 < 16; fi8++ {
        v4 = 1.5
        if (!(int64(v1) < (v5 & 42))) {
            var v9 bool = (!(v1 >= (v1 ^ (v1 >> ((42) & 63)))))
            for fi10 := 0; fi10 < 6; fi10++ {
                if v9 {
                    if (!(-2.25 > 1.5)) {
                        var v11 uint32 = v7
                        if v9 {
                            v12 := GenS1{F0: 23}
                            _ = v12
                        }
                        _ = v11
                    }
                    v13 := GenS2{F0: 255}
                    if v9 {
                        var v14 string = "hello"
                        var v15 int16 = v1
                        switch v14 {
                        case "\n\t":
                            v1 = 42
                        case "foo bar":
                            fmt.Println(v15)
                        default:
                            v16 := GenS1{F0: 0}
                            _ = v16
                        }
                        _ = v14
                        _ = v15
                    }
                    _ = v13
                }
                _ = fi10
            }
            fmt.Println(v1, v1, v5)
            _ = v9
        } else {
            var v17 uint = 255
            _ = v17
        }
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 uint32 = v7
            _ = fi18
            _ = v19
        }
        _ = fi8
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        if false {
            if false {
                v21 := []int{}
                if false {
                    fmt.Println(v6, fi20, v6)
                    switch "a" {
                    case "a":
                        fmt.Println(v7)
                    case "\n\t":
                        v7 = 0
                    case "":
                        fmt.Println(v5, fi20, fi20)
                    default:
                        for fi22 := 0; fi22 < 16; fi22++ {
                            if false {
                                var v23 int8 = 57
                                v2[("a" + "a")] = true
                                _ = v23
                            }
                            _ = fi22
                        }
                    }
                    var v24 uint16 = 18
                    _ = v24
                } else {
                    v1 -= v1
                    if true {
                        if ((v7 > (^(v7))) && ((("foo bar" + "foo bar") + "") >= (((((("a" + "a") + "hello") + ("a" + "hello")) + "\n\t") + ("foo bar" + ("\n\t" + ("" + "foo bar")))) + ""))) {
                            var v25 uint32 = (^(v7))
                            v26 := GenS2{F0: 255}
                            if (true || (1e10 != float64(((v4 + (v4 / (v4 - v4))) * float32((1e10 / ((0.0 / ((float64(v4) * 0.1) - 0.0)) * 1.5))))))) {
                                var v27 int8 = 0
                                _ = v27
                            } else {
                                for fi28 := 0; fi28 < 5; fi28++ {
                                    switch "a" {
                                    case "\n\t":
                                        switch "a" {
                                        case "hello":
                                            var v29 uint = 65535
                                            _ = v29
                                        }
                                    case "hello":
                                        v30 := GenS3{F0: 0, F1: 952, F2: v5}
                                        _ = v30
                                    default:
                                        v5 = (v5 - v5)
                                    }
                                    if (1 == -9223372036854775808) {
                                        v25 -= (v7 << ((4294967295) & 63))
                                        var v31 uint32 = (v7 ^ v7)
                                        _ = v31
                                    }
                                    _ = fi28
                                }
                            }
                            _ = v25
                            _ = v26
                        }
                        for fi32 := 0; fi32 < 14; fi32++ {
                            fmt.Println(fi20, v4, v1)
                            _ = fi32
                        }
                    } else {
                        var v33 int = 651
                        for fi34 := 0; fi34 < 2; fi34++ {
                            switch "a" {
                            case "foo bar":
                                var v35 uint16 = 65535
                                _ = v35
                            case "a":
                                var v36 uint32 = (v7 ^ (v7 >> ((0) & 63)))
                                _ = v36
                            case "":
                                var v37 int32 = 42
                                _ = v37
                            }
                            _ = fi34
                        }
                        v38 := GenS1{F0: 0}
                        _ = v33
                        _ = v38
                    }
                    v7 = v7
                }
                var v39 int = 2147483647
                _ = v21
                _ = v39
            } else {
                v6 += 4294967295
                var v40 uint16 = 0
                _ = v40
            }
        } else {
            var v41 int = 127
            _ = v41
        }
        var v42 int8 = int8(v6)
        var v43 bool = true
        _ = fi20
        _ = v42
        _ = v43
    }
    if (!(-2147483648 > -1)) {
        for fi44 := 0; fi44 < 10; fi44++ {
            fmt.Println(v5, v6, v4)
            for fi45 := 0; fi45 < 14; fi45++ {
                v5 = v5
                _ = fi45
            }
            var v46 float32 = v4
            _ = fi44
            _ = v46
        }
    }
    if false {
        fmt.Println(v5, v4)
        var v47 int32 = int32(v1)
        _ = v47
    }
    fmt.Println(v1, v6, v4)
    if (v6 <= v6) {
        fmt.Println(v1, v6)
        if (52 == 877) {
            v48 := GenS4{F0: "a", F1: 237}
            v4 -= float32((1.5 + 3.14159))
            _ = v48
        }
    }
    v49 := []int{32767, 32768, 127, int(v5)}
    for fi50 := 0; fi50 < 7; fi50++ {
        if (true || (2147483648 != 65535)) {
            var v51 string = "hello"
            v1 -= 0
            v6 = 0
            _ = v51
        } else {
            var v52 uint8 = 194
            _ = len(v2)
            v7 += v7
            _ = v52
        }
        v6 -= 2
        _ = fi50
    }
    _ = len(v49)
    v53 := v3[(v7 & v7)]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v53
    _ = v6
    _ = v7
}
