package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16 }
func genFn0(p0 int16, p1 uint16, p2 uint) int8 {
    p2 = 4294967295
    var v6 uint = (p2 >> ((158) & 63))
    switch "\n\t" {
    case "hello":
        var v7 float64 = 0.0
        _ = v7
    case "a":
        v8 := []int{}
        _ = v8
    case "":
        p0 += p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 53
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 0
    v3 := make(map[uint32]int)
    var v4 int8 = 100
    var v5 uint = 65535
    switch 127 {
    case 1:
        switch "foo bar" {
        case "a":
            v9 := []int{-388}
            _ = v9
        case "foo bar":
            for fi10 := 0; fi10 < 2; fi10++ {
                v2 -= (v5 & 626)
                if false {
                    for fi11 := 0; fi11 < 2; fi11++ {
                        switch "\n\t" {
                        case "a":
                            if true {
                                v12 := GenS1{F0: 0}
                                _ = v12
                            }
                        default:
                            var v13 string = "foo bar"
                            _ = v13
                        }
                        v4 = -2
                        _ = fi11
                    }
                } else {
                    var v14 float32 = (-2.25 - 3.14159)
                    for fi15 := 0; fi15 < 19; fi15++ {
                        var v16 uint8 = 127
                        _ = fi15
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 16; fi17++ {
                        var v18 int64 = -9223372036854775808
                        v19 := GenS1{F0: 2147483648}
                        v4 = 0
                        var v20 float32 = float32((1.5 + 1.5))
                        switch "hello" {
                        case "\n\t":
                            v21 := []int{}
                            _ = v21
                        case "":
                            switch "a" {
                            case "foo bar":
                                v22 := GenS1{F0: 42}
                                _ = v22
                            case "hello":
                                if true {
                                    v23 := GenS2{F0: 127}
                                    v24 := GenS1{F0: 1000}
                                    _ = v23
                                    _ = v24
                                } else {
                                    switch fi17 {
                                    case 32768:
                                        fmt.Println(v20, v20, v2)
                                    case -9223372036854775808:
                                        v25 := []int{127}
                                        _ = v25
                                    default:
                                        if false {
                                            var v26 string = ("\n\t" + "a")
                                            var v27 float32 = v14
                                            _ = v26
                                            _ = v27
                                        }
                                    }
                                }
                            case "":
                                fmt.Println(v2, v2)
                            default:
                                _ = len(v1)
                            }
                        }
                        var v28 int = (fi10 - fi10)
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                        _ = v28
                    }
                    _ = v14
                }
                _ = fi10
            }
        default:
            switch "a" {
            case "":
                var v29 int64 = 9223372036854775807
                _ = v29
            case "foo bar":
                delete(v3, 65535)
            case "a":
                v30 := GenS1{F0: 725}
                _ = v30
            }
        }
    case 9223372036854775807:
        var v31 float32 = 0.0
        _ = v31
    case -2147483648:
        if false {
            var v32 uint = (v5 % ((v5) | 1))
            _ = v32
        }
    }
    v33 := GenS2{F0: 60}
    switch ("a" + "\n\t") {
    case "a":
        var v34 int32 = int32(v2)
        _ = v34
    default:
        var v35 uint = v5
        _ = v35
    }
    v36 := GenS3{F0: "a"}
    fmt.Println(v2, v4, v2)
    v4 = v4
    v2 -= (v5 + 4294967295)
    var v37 float64 = 3.14159
    v2 = (^(v2))
    fmt.Println(v37, v5, v2)
    switch ("foo bar" + "") {
    case "\n\t":
        switch 2147483647 {
        case 42:
            var v38 int8 = (v4 ^ 92)
            _ = v38
        default:
            var v39 float64 = v37
            _ = v39
        }
    case "a":
        var v40 bool = (v2 >= v5)
        _ = v40
    }
    fmt.Println(v37, v4, v37)
    v41 := []int{269, 12, 128}
    v42 := GenS2{F0: 105}
    v43 := GenS1{F0: 42}
    for fi44 := 0; fi44 < 3; fi44++ {
        for fi45 := 0; fi45 < 3; fi45++ {
            v37 += (-2.25 * v37)
            v46 := GenS4{F0: -113}
            _ = fi45
            _ = v46
        }
        var v47 int8 = (v4 - v4)
        _ = fi44
        _ = v47
    }
    for fi48 := 0; fi48 < 16; fi48++ {
        for fi49 := 0; fi49 < 14; fi49++ {
            var v50 int8 = v4
            _ = fi49
            _ = v50
        }
        _ = fi48
    }
    if (!true) {
        var v51 uint = (v5 ^ 340)
        v37 -= 3.14159
        var v52 int8 = 127
        _ = v51
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
