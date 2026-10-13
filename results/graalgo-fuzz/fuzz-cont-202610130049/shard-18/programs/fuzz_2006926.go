package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 bool }
func genFn0(p0 string, p1 uint8) int32 {
    v5 := []int{-9223372036854775808}
    switch p0 {
    case "hello":
        if (-32768 >= 86) {
            var v6 string = p0
            _ = v6
        } else {
            if ((p1 & (p1 / (((p1 << ((65535) & 63))) | 1))) == p1) {
                for fi7 := 0; fi7 < 16; fi7++ {
                    fmt.Println(p1, p1)
                    v8 := GenS1{F0: true}
                    var v9 uint16 = 2
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            }
        }
    case "foo bar":
        for fi10 := 0; fi10 < 1; fi10++ {
            for fi11 := 0; fi11 < 3; fi11++ {
                v12 := GenS2{F0: -32768, F1: -1}
                var v13 float32 = (0.1 * (1.5 - 0.1))
                var v14 int32 = -1
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            fmt.Println(p0)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi10
        }
    default:
        p1 -= p1
    }
    _ = p0
    _ = p1
    _ = v5
    return -2147483648
}
func genFn1() bool {
    var v15 int64 = -56
    v16 := []int{32768, -1, 2147483647}
    _ = v15
    _ = v16
    return (false || (-2.25 <= float32(float64(1.5))))
}

func main() {
    var v1 bool = true
    var v2 int = -400
    var v3 uint16 = 98
    var v4 int = (v2 | v2)
    if (!v1) {
        fmt.Println(v4, v4, v4)
        v17 := GenS3{F0: (("\n\t" + "\n\t") + "foo bar"), F1: 128}
        var v18 uint32 = 65535
        _ = v17
        _ = v18
    }
    v4 = (v2 % ((42) | 1))
    switch ("a" + "hello") {
    case "\n\t":
        for fi19 := 0; fi19 < 17; fi19++ {
            v20 := []int{v2, (-(v4))}
            if (255 <= 0) {
                fmt.Println(v4, v1, fi19)
                for fi21 := 0; fi21 < 11; fi21++ {
                    switch ("hello" + "") {
                    case "hello":
                        if ((-(v2)) >= (v2 - 4294967295)) {
                            _ = len(v20)
                            v4 -= fi21
                            switch ((("\n\t" + "hello") + ("foo bar" + "\n\t")) + "\n\t") {
                            case "\n\t":
                                var v22 uint = 1
                                _ = v22
                            case "a":
                                v23 := GenS3{F0: "a", F1: 2}
                                _ = v23
                            }
                        } else {
                            var v24 int32 = 1
                            v3 += v3
                            var v25 string = ((("" + "hello") + "") + (("a" + "hello") + "a"))
                            _ = v24
                            _ = v25
                        }
                    }
                    var v26 bool = (-1 > 9223372036854775807)
                    if (127 == 255) {
                        fmt.Println(fi19)
                        if len(v20) > 0 {
                            v20[0] = v4
                        }
                        if v1 {
                            var v27 string = (("\n\t" + "foo bar") + "a")
                            _ = v27
                        }
                    } else {
                        var v28 uint16 = (^(v3))
                        _ = v28
                    }
                    if v1 {
                        var v29 float32 = -2.25
                        _ = v29
                    }
                    _ = fi21
                    _ = v26
                }
            } else {
                if len(v20) > 0 {
                    v20[0] = 9223372036854775807
                }
                v30 := GenS4{F0: 42, F1: 255, F2: 216}
                _ = v30
            }
            for fi31 := 0; fi31 < 19; fi31++ {
                switch "" {
                case "hello":
                    v32 := GenS4{F0: 0, F1: 7, F2: 1}
                    fmt.Println(v3, v3)
                    fmt.Println(v2, fi31)
                    fmt.Println(fi31, v3)
                    _ = v32
                case "foo bar":
                    v33 := GenS3{F0: "\n\t", F1: 183}
                    _ = v33
                case "":
                    fmt.Println(fi31)
                }
                var v34 uint32 = 950
                _ = fi31
                _ = v34
            }
            _ = fi19
            _ = v20
        }
    case "a":
        switch 128 {
        case 128:
            v35 := []int{(^(v2)), v4, (v4 / (((v4 - 4294967295)) | 1)), v2}
            _ = v35
        case 457:
            v1 = v1
        }
    case "":
        var v36 int32 = 42
        _ = v36
    }
    var v37 int16 = 42
    v4 += (v2 % (((v4 | -15)) | 1))
    v38 := []int{v4, (^(v2)), 233}
    v39 := GenS5{F0: -74, F1: ("foo bar" + "a")}
    switch (v2 * (v2 % (((v4 | v2)) | 1))) {
    case 32767:
        var v40 uint64 = 0
        _ = v40
    default:
        if (false && v1) {
            v41 := GenS6{F0: ((2 <= 1) && (0 > 4294967295))}
            _ = v41
        }
    }
    switch "foo bar" {
    case "":
        var v42 int16 = v37
        _ = v42
    default:
        var v43 int64 = -1
        _ = v43
    }
    var v44 uint16 = v3
    v2 -= v4
    if ((((2147483648 != 2147483647) || (((v3 << ((2) & 63)) < v44) && (234 == 2))) && (352 >= 655)) || false) {
        v4 = 650
        var v45 int8 = 7
        var v46 int8 = (v45 & 42)
        _ = v45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v44
}
