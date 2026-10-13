package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int16 }
func genFn0(p0 int64, p1 uint16, p2 bool) string {
    var v4 int = 127
    v5 := GenS1{F0: 100, F1: 1, F2: 546}
    var v6 string = "foo bar"
    if p2 {
        v7 := []int{}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return (("foo bar" + v6) + ((v6 + v6) + v6))
}
func genFn1(p0 int8) int64 {
    var v8 uint = 42
    v8 -= v8
    _ = p0
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    if (true || false) {
        var v9 string = "\n\t"
        _ = v9
    } else {
        v10 := GenS1{F0: -42, F1: 127, F2: 0}
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 uint64 = 2
            switch ("\n\t" + "hello") {
            case "foo bar":
                if (!true) {
                    var v13 float32 = 1e10
                    v13 = ((v13 / v13) * 0.0)
                    _ = v13
                } else {
                    var v14 bool = false
                    _ = v14
                }
            case "hello":
                var v15 float32 = (0.1 - (1e10 + 1e10))
                _ = v15
            case "a":
                var v16 bool = false
                _ = v16
            }
            _ = fi11
            _ = v12
        }
        if (0.1 != 0.0) {
            if false {
                var v17 uint8 = 42
                if ((1 >= 18446744073709551615) || false) {
                    var v18 uint = 2
                    switch ("" + "foo bar") {
                    case "foo bar":
                        var v19 int16 = 32767
                        _ = v19
                    case "\n\t":
                        var v20 uint8 = 42
                        _ = v20
                    case "hello":
                        v18 -= v18
                    default:
                        if (true || true) {
                            v17 -= 162
                            var v21 bool = (true && false)
                            v22 := []int{}
                            _ = v21
                            _ = v22
                        } else {
                            v23 := []int{0, 32767}
                            if (!true) {
                                v24 := v1[4294967295]
                                _ = v24
                            } else {
                                v25 := GenS2{F0: -1}
                                var v26 uint64 = 1
                                _ = v25
                                _ = v26
                            }
                            v27 := GenS3{F0: (("a" + ("foo bar" + ("hello" + "a"))) + ""), F1: -32768}
                            _ = v23
                            _ = v27
                        }
                    }
                    _ = v18
                } else {
                    v17 = v17
                    v28 := []int{2147483647, -517}
                    _ = v28
                }
                _ = v17
            } else {
                var v29 int16 = -97
                v30 := v1[9223372036854775807]
                fmt.Println(v29, v29)
                _ = v29
                _ = v30
            }
        } else {
            for fi31 := 0; fi31 < 13; fi31++ {
                var v32 uint64 = 0
                _ = fi31
                _ = v32
            }
            if true {
                switch 153 {
                case 860:
                    var v33 float64 = 1e10
                    _ = v33
                case 4294967295:
                    var v34 uint32 = 2147483648
                    _ = v34
                case -721:
                    var v35 string = ("" + "a")
                    _ = v35
                default:
                    switch (("hello" + "foo bar") + ("foo bar" + "")) {
                    case "":
                        v36 := []int{}
                        _ = v36
                    case "a":
                        v37 := GenS2{F0: 0}
                        _ = v37
                    case "\n\t":
                        switch "hello" {
                        case "\n\t":
                            v38 := v3[0]
                            _ = v38
                        case "a":
                            var v39 int16 = 888
                            _ = v39
                        default:
                            var v40 bool = (false || (!true))
                            _ = v40
                        }
                    }
                }
                var v41 int8 = 127
                _ = v41
            }
            v1[9223372036854775807] = "foo bar"
        }
        _ = v10
    }
    if false {
        v42 := GenS2{F0: 1}
        _ = v42
    }
    switch ("\n\t" + ((("" + ("\n\t" + "hello")) + "\n\t") + (("" + "foo bar") + "a"))) {
    case "foo bar":
        var v43 int8 = -58
        _ = v43
    default:
        var v44 int8 = 21
        _ = v44
    }
    switch "" {
    case "foo bar":
        var v45 uint = 255
        _ = v45
    case "\n\t":
        var v46 uint8 = 0
        _ = v46
    case "hello":
        var v47 int16 = -32768
        _ = v47
    }
    var v48 int64 = 9223372036854775807
    v49 := []int{32767, int(v48)}
    if (!true) {
        v1[127] = ("foo bar" + "a")
    } else {
        v48 = (v48 >> ((255) & 63))
        v48 = v48
        for fi50 := 0; fi50 < 1; fi50++ {
            fmt.Println(v48)
            _ = fi50
        }
    }
    var v51 uint8 = 0
    switch "\n\t" {
    case "":
        var v52 uint16 = 42
        _ = v52
    default:
        var v53 uint32 = 0
        _ = v53
    }
    fmt.Println(v48)
    _ = len(v2)
    var v54 int = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v51
    _ = v54
}
