package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 int16; F2 int8 }
type GenS4 struct { F0 int8; F1 int64 }

func main() {
    var v1 float64 = 1e10
    var v2 uint8 = 127
    var v3 int32 = 2147483647
    v4 := make(map[string]int)
    switch ("\n\t" + "\n\t") {
    case "a":
        var v5 int64 = 801
        _ = v5
    case "":
        var v6 float32 = (-2.25 + 0.0)
        _ = v6
    case "\n\t":
        var v7 uint32 = 87
        _ = v7
    default:
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := []int{(fi8 - (^(fi8))), fi8, fi8}
            _ = fi8
            _ = v9
        }
    }
    if true {
        if false {
            v10 := []int{42, 127, -286}
            switch "foo bar" {
            case "":
                v4["a"] = 4294967295
            case "foo bar":
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            case "\n\t":
                fmt.Println(v1, v3)
            default:
                switch ("foo bar" + "") {
                case "a":
                    v1 -= v1
                }
            }
            var v11 uint64 = 1
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 4; fi12++ {
            if true {
                _ = len(v4)
            } else {
                if (false || false) {
                    v3 += 2147483647
                }
                for fi13 := 0; fi13 < 1; fi13++ {
                    v14 := []int{}
                    var v15 float32 = 1e10
                    v1 += v1
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
            var v16 string = (("\n\t" + ("\n\t" + ("\n\t" + ("foo bar" + "foo bar")))) + "")
            switch "hello" {
            case "foo bar":
                for fi17 := 0; fi17 < 12; fi17++ {
                    v2 = (v2 - (v2 >> ((18446744073709551615) & 63)))
                    v1 = v1
                    _ = fi17
                }
            case "":
                var v18 uint32 = 2147483648
                _ = v18
            case "hello":
                switch fi12 {
                case 9223372036854775807:
                    v3 -= v3
                case -986:
                    for fi19 := 0; fi19 < 20; fi19++ {
                        v20 := []int{(fi19 % ((fi12) | 1)), (fi12 + fi12)}
                        v21 := GenS1{F0: 42}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                case 102:
                    if (452 >= 2) {
                        v22 := GenS1{F0: 4294967295}
                        if (((!true) || true) || (!true)) {
                            fmt.Println(fi12, v2)
                            fmt.Println(fi12)
                            if (v3 < v3) {
                                for fi23 := 0; fi23 < 3; fi23++ {
                                    var v24 int8 = 90
                                    v3 = v3
                                    v25 := GenS2{F0: 42}
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                }
                                fmt.Println(v1, v16)
                            } else {
                                if (true && (!false)) {
                                    v26 := []int{}
                                    _ = v26
                                }
                            }
                        } else {
                            v27 := GenS1{F0: 0}
                            fmt.Println(v2, v3, v3)
                            _ = v27
                        }
                        v28 := []int{fi12, 787, (fi12 | 4294967295)}
                        _ = v22
                        _ = v28
                    }
                }
            default:
                switch fi12 {
                case -9223372036854775808:
                    v29 := []int{(fi12 + (fi12 & -1))}
                    _ = v29
                default:
                    v16 = ((v16 + (((v16 + ("\n\t" + "")) + v16) + v16)) + (v16 + "a"))
                }
            }
            _ = fi12
            _ = v16
        }
    } else {
        var v30 uint = 42
        var v31 uint32 = 688
        switch "hello" {
        case "foo bar":
            v30 += (^(v30))
        case "":
            fmt.Println(v1, v31, v2)
        }
        _ = v30
        _ = v31
    }
    v2 = (v2 - (v2 / ((18) | 1)))
    for fi32 := 0; fi32 < 11; fi32++ {
        if true {
            var v33 int64 = -116
            v4[(("\n\t" + "") + "a")] = fi32
            v34 := v4["hello"]
            _ = v33
            _ = v34
        }
        v35 := []int{fi32}
        _ = fi32
        _ = v35
    }
    switch ("foo bar" + ("foo bar" + (("" + "hello") + "\n\t"))) {
    case "hello":
        switch 0 {
        case 42:
            v2 -= 2
        case 9223372036854775807:
            if false {
                v1 += 3.14159
                if (float32(v1) < (3.14159 - 3.14159)) {
                    v36 := GenS2{F0: 42}
                    _ = v36
                }
                fmt.Println(v2, v2, v3)
            } else {
                v37 := v4["hello"]
                _ = v37
            }
        case -801:
            v38 := GenS1{F0: 4294967295}
            _ = v38
        }
    case "foo bar":
        v39 := GenS3{F0: -9223372036854775808, F1: 0, F2: 124}
        _ = v39
    }
    v40 := GenS4{F0: 100, F1: -463}
    var v41 int64 = 42
    for fi42 := 0; fi42 < 1; fi42++ {
        switch (fi42 >> ((1) & 63)) {
        case 0:
            for fi43 := 0; fi43 < 8; fi43++ {
                switch ("foo bar" + ("\n\t" + "")) {
                case "hello":
                    var v44 uint32 = 957
                    _ = v44
                }
                fmt.Println(fi43, fi43, v2)
                _ = fi43
            }
        case 128:
            delete(v4, "a")
        }
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
