package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint32; F1 uint }
func genFn0() int32 {
    v7 := []int{128, 127}
    var v8 uint16 = 255
    var v9 uint64 = 0
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 int8 = 39
    var v2 uint = 255
    var v3 int = 875
    v4 := make(map[string]string)
    var v5 uint64 = 1
    var v6 uint = 481
    v10 := []int{(v3 / (((v3 << ((v2) & 63))) | 1))}
    for fi11 := 0; fi11 < 6; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            switch (fi12 >> ((628) & 63)) {
            case 127:
                if (uint8(v5) <= 43) {
                    var v13 uint32 = 615
                    v5 += (v5 * 4294967295)
                    if false {
                        if false {
                            v14 := GenS1{F0: genFn0()}
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            _ = v14
                        } else {
                            v1 = v1
                            var v15 uint32 = 2147483648
                            v16 := GenS1{F0: 42}
                            switch 32767 {
                            case 2147483647:
                                var v17 int64 = 0
                                _ = v17
                            }
                            v15 -= v15
                            fmt.Println(v2, v15)
                            var v18 uint16 = 697
                            v19 := v4[""]
                            _ = v15
                            _ = v16
                            _ = v18
                            _ = v19
                        }
                        fmt.Println(v3, fi12, v1)
                    } else {
                        if true {
                            v20 := GenS1{F0: 42}
                            var v21 uint64 = 42
                            _ = v20
                            _ = v21
                        }
                        v1 += 22
                    }
                    _ = v13
                }
            }
            switch (fi12 / (((v3 | 32768)) | 1)) {
            case 127:
                var v22 string = "foo bar"
                _ = v22
            case 32768:
                v10 = append(v10, fi12)
            default:
                v3 -= (v3 >> ((4294967295) & 63))
            }
            if (-872 != -306) {
                if ((-128 > 36) || ((-32768 >= -1) && true)) {
                    if (("hello" <= "\n\t") && (v6 == v6)) {
                        switch ("" + (("hello" + "") + "")) {
                        case "":
                            v23 := GenS1{F0: -2147483648}
                            _ = v23
                        default:
                            var v24 int32 = genFn0()
                            v10 = append(v10, 9223372036854775807)
                            _ = v24
                        }
                        v25 := GenS2{F0: (255 == 237)}
                        var v26 uint16 = uint16(v2)
                        var v27 int = fi11
                        var v28 uint64 = v5
                        _ = v25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    var v29 float64 = -2.25
                    var v30 int16 = 1
                    var v31 int64 = -149
                    v32 := GenS2{F0: (42 != 2147483648)}
                    _ = v29
                    _ = v30
                    _ = v31
                    _ = v32
                } else {
                    delete(v4, ("a" + "\n\t"))
                    var v33 uint32 = 0
                    v3 -= fi12
                    _ = v33
                }
            }
            _ = fi12
        }
        v34 := GenS1{F0: -1}
        _ = fi11
        _ = v34
    }
    v2 = 237
    if (1.5 != -2.25) {
        switch ("a" + "hello") {
        case "a":
            if true {
                var v35 uint16 = 0
                _ = v35
            } else {
                fmt.Println(v3, v5, v6)
                var v36 string = ""
                _ = v36
            }
        }
        var v37 uint32 = 2147483648
        _ = v37
    } else {
        if (!false) {
            if true {
                v1 = (v1 / ((v1) | 1))
                v38 := GenS1{F0: 2147483647}
                var v39 int = v3
                _ = v38
                _ = v39
            } else {
                fmt.Println(v5, v3)
                fmt.Println(v1, v3)
            }
        } else {
            var v40 bool = (!true)
            switch "foo bar" {
            case "":
                var v41 int32 = -1
                _ = v41
            case "a":
                for fi42 := 0; fi42 < 9; fi42++ {
                    switch "foo bar" {
                    case "a":
                        var v43 uint = (^(v2))
                        _ = v43
                    }
                    fmt.Println(v6)
                    _ = fi42
                }
            }
            _ = v40
        }
    }
    for fi44 := 0; fi44 < 7; fi44++ {
        v45 := GenS3{F0: 0, F1: ""}
        v4["foo bar"] = ("a" + (("" + "\n\t") + ""))
        _ = fi44
        _ = v45
    }
    delete(v4, "foo bar")
    if (!false) {
        for fi46 := 0; fi46 < 12; fi46++ {
            fmt.Println(v6, v1, v1)
            v47 := GenS3{F0: 514, F1: ""}
            _ = fi46
            _ = v47
        }
    } else {
        var v48 uint8 = 245
        _ = len(v10)
        _ = v48
    }
    v49 := GenS1{F0: genFn0()}
    switch "" {
    case "hello":
        v5 = (^(v5))
    case "":
        for fi50 := 0; fi50 < 13; fi50++ {
            v51 := GenS4{F0: ((("hello" + "hello") + (("\n\t" + (("\n\t" + "a") + ("" + "\n\t"))) + "")) + "hello"), F1: v6}
            v3 -= (^(v3))
            v52 := GenS3{F0: 42, F1: "hello"}
            _ = fi50
            _ = v51
            _ = v52
        }
    }
    var v53 uint64 = v5
    v54 := GenS5{F0: 2, F1: v2}
    v1 += (v1 / ((0) | 1))
    for fi55 := 0; fi55 < 20; fi55++ {
        var v56 int16 = -281
        switch "hello" {
        case "":
            if false {
                var v57 uint = 913
                _ = v57
            } else {
                var v58 bool = false
                _ = v58
            }
        case "foo bar":
            switch -86 {
            case 103:
                var v59 bool = (!true)
                _ = v59
            case 42:
                for fi60 := 0; fi60 < 16; fi60++ {
                    if true {
                        var v61 bool = ((fi60 / (((v3 & fi60)) | 1)) <= (^(fi60)))
                        _ = v61
                    } else {
                        var v62 int = v3
                        _ = v62
                    }
                    v56 = 1000
                    _ = fi60
                }
            case 531:
                v56 -= (v56 & v56)
            }
        case "a":
            switch "" {
            case "a":
                v6 -= 4294967295
            case "\n\t":
                var v63 float32 = float32(float64((1.5 - (1e10 * 0.0))))
                _ = v63
            }
        }
        _ = fi55
        _ = v56
    }
    var v64 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v53
    _ = v54
    _ = v6
    _ = v64
}
