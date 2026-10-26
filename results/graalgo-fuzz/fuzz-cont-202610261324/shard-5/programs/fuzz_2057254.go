package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"\n\t"}
    v3 := make(map[uint32]bool)
    var v4 float32 = 1.5
    v4 = (v4 * (-2.25 * v4))
    fmt.Println(v4)
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 bool = false
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 1, F1: 0, F2: 0}
    fmt.Println(v4)
    v4 = float32(0.0)
    v8 := v3[65535]
    fmt.Println(v8, v4)
    switch -1 {
    case 42:
        v8 = false
    case 763:
        var v9 string = ""
        _ = v9
    }
    if v8 {
        var v10 int64 = 496
        delete(v3, 42)
        var v11 uint = 4294967295
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 14; fi12++ {
            switch "" {
            case "\n\t":
                if v8 {
                    fmt.Println(v4, fi12, fi12)
                    var v13 bool = (65535 <= 255)
                    var v14 int32 = 146
                    _ = v13
                    _ = v14
                } else {
                    var v15 float64 = 1.5
                    if true {
                        switch "a" {
                        case "":
                            v15 = v15
                        case "hello":
                            for fi16 := 0; fi16 < 8; fi16++ {
                                v8 = (0 < 0)
                                fmt.Println(fi16, v4)
                                switch ("hello" + (((("" + "\n\t") + ("a" + "hello")) + "foo bar") + "hello")) {
                                case "hello":
                                    var v17 int32 = 42
                                    _ = v17
                                case "\n\t":
                                    var v18 int16 = 1
                                    _ = v18
                                case "foo bar":
                                    if false {
                                        v19 := v1[4294967295]
                                        _ = v19
                                    } else {
                                        switch 1 {
                                        case 9223372036854775807:
                                            if ((!(9223372036854775807 != 0)) || true) {
                                                var v20 bool = (133 > 255)
                                                _ = len(v1)
                                                var v21 uint16 = 255
                                                _ = v20
                                                _ = v21
                                            }
                                        case 1:
                                            v2 = append(v2, ("hello" + (("" + "\n\t") + "foo bar")))
                                            if true {
                                                var v22 int = (fi16 | (^(fi16)))
                                                var v23 uint = 42
                                                var v24 string = (("\n\t" + ((("" + "hello") + ("\n\t" + ("a" + "foo bar"))) + "hello")) + "a")
                                                _ = v22
                                                _ = v23
                                                _ = v24
                                            } else {
                                                v25 := GenS2{F0: 155}
                                                v8 = v8
                                                if false {
                                                    var v26 uint16 = 360
                                                    if len(v2) > 0 {
                                                        v2[0] = "foo bar"
                                                    }
                                                    if false {
                                                        fmt.Println(v8)
                                                    } else {
                                                        v8 = v8
                                                        v15 += v15
                                                        v27 := GenS3{F0: -49}
                                                        _ = v27
                                                    }
                                                    _ = v26
                                                }
                                                _ = v25
                                            }
                                        default:
                                            if ((false || (((("foo bar" + "foo bar") + ("hello" + "foo bar")) + ((("a" + "hello") + "hello") + "hello")) <= "")) && true) {
                                                fmt.Println(v4, fi16, v15)
                                                var v28 uint64 = 42
                                                fmt.Println(v15, v4)
                                                _ = v28
                                            }
                                        }
                                        var v29 uint16 = 658
                                        v30 := GenS3{F0: 100}
                                        _ = v29
                                        _ = v30
                                    }
                                }
                                _ = fi16
                            }
                        }
                    }
                    _ = v15
                }
            case "":
                if v8 {
                    if false {
                        for fi31 := 0; fi31 < 1; fi31++ {
                            v8 = v8
                            v8 = true
                            var v32 uint32 = 4294967295
                            _ = fi31
                            _ = v32
                        }
                        fmt.Println(fi12, fi12)
                        var v33 int32 = -2147483648
                        _ = v33
                    } else {
                        if v8 {
                            v2 = append(v2, "hello")
                        }
                        _ = len(v2)
                    }
                    switch ("foo bar" + ("hello" + "\n\t")) {
                    case "a":
                        var v34 float32 = 1e10
                        _ = v34
                    default:
                        v8 = v8
                    }
                } else {
                    switch fi12 {
                    case 710:
                        switch "a" {
                        case "a":
                            for fi35 := 0; fi35 < 17; fi35++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v36 uint16 = 128
                                v37 := GenS3{F0: 54}
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                        default:
                            switch "a" {
                            case "":
                                for fi38 := 0; fi38 < 14; fi38++ {
                                    var v39 uint8 = 0
                                    _ = len(v2)
                                    v40 := GenS4{F0: -670}
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                            case "foo bar":
                                v4 -= (v4 / v4)
                            case "hello":
                                v8 = true
                            }
                        }
                    default:
                        v8 = (580 < 316)
                    }
                    v8 = (0 <= -1)
                }
            case "hello":
                _ = len(v2)
            }
            switch (("hello" + "") + (("foo bar" + "a") + "hello")) {
            case "hello":
                fmt.Println(fi12, v4, v4)
            case "foo bar":
                fmt.Println(fi12)
            case "":
                v4 += v4
            }
            _ = fi12
        }
    }
    if v8 {
        var v41 uint32 = 65535
        _ = v41
    } else {
        v42 := GenS2{F0: 2}
        _ = v42
    }
    for fi43 := 0; fi43 < 18; fi43++ {
        var v44 uint = 255
        _ = fi43
        _ = v44
    }
    v4 = v4
    v8 = (false && ((v4 + v4) <= v4))
    v4 -= -2.25
    _ = len(v2)
    v8 = ((!v8) && true)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
