package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 string; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int; F1 bool; F2 int8 }

func main() {
    var v1 float64 = -2.25
    var v2 bool = true
    v3 := make(map[string]int)
    var v4 uint = 2
    var v5 float32 = 0.1
    switch "" {
    case "\n\t":
        fmt.Println(v1, v4)
    default:
        var v6 float64 = 0.1
        _ = v6
    }
    if (9223372036854775807 != 9223372036854775807) {
        v2 = v2
        v7 := GenS1{F0: 2147483647, F1: "hello"}
        if true {
            v1 += float64((v5 - (v5 + v5)))
        } else {
            var v8 int16 = 861
            _ = v8
        }
        _ = v7
    }
    var v9 uint64 = 4294967295
    switch -9223372036854775808 {
    case 42:
        v9 = (v9 + v9)
    }
    if v2 {
        var v10 bool = v2
        v11 := GenS1{F0: -2147483648, F1: ("foo bar" + "foo bar")}
        v12 := GenS1{F0: -2147483648, F1: ("\n\t" + ("\n\t" + ("\n\t" + "\n\t")))}
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := []int{0, -9223372036854775808, 885, 32767}
    var v14 float32 = v5
    delete(v3, ("" + (("a" + "\n\t") + ("hello" + ("hello" + "\n\t")))))
    switch 522 {
    case 32767:
        for fi15 := 0; fi15 < 3; fi15++ {
            switch (-(fi15)) {
            case 1:
                var v16 float32 = (v5 - 0.0)
                _ = v16
            default:
                if v2 {
                    var v17 int = fi15
                    fmt.Println(v14)
                    _ = v17
                }
            }
            for fi18 := 0; fi18 < 2; fi18++ {
                if v2 {
                    if ((true && (!v2)) && v2) {
                        if v2 {
                            switch (fi15 >> ((v4) & 63)) {
                            case 717:
                                var v19 int32 = 2147483647
                                _ = v19
                            case 9223372036854775807:
                                v20 := GenS1{F0: 0, F1: ("" + "\n\t")}
                                _ = v20
                            case 1:
                                switch "\n\t" {
                                case "hello":
                                    if (42 <= 69) {
                                        var v21 uint32 = 919
                                        _ = v21
                                    }
                                    if (0 >= 128) {
                                        if true {
                                            switch "" {
                                            case "foo bar":
                                                v2 = v2
                                            case "hello":
                                                fmt.Println(fi18, v9, v5)
                                            default:
                                                v4 = (v4 & (v4 << ((v4) & 63)))
                                            }
                                        }
                                        v14 -= -2.25
                                        if (fi15 == fi18) {
                                            fmt.Println(fi18)
                                            var v22 uint8 = 243
                                            switch (fi18 - (-(fi18))) {
                                            case 197:
                                                v3["foo bar"] = (-(fi18))
                                            case 2147483647:
                                                var v23 int32 = 42
                                                _ = v23
                                            case -9223372036854775808:
                                                v24 := GenS2{F0: ((42 < -199) && (65535 != 4294967295))}
                                                _ = v24
                                            default:
                                                v25 := GenS2{F0: false}
                                                _ = v25
                                            }
                                            if (true || v2) {
                                                var v26 int = fi15
                                                if false {
                                                    v2 = v2
                                                    v27 := GenS3{F0: 298, F1: "", F2: 1}
                                                    v28 := v3["a"]
                                                    _ = v27
                                                    _ = v28
                                                }
                                                v14 = 1.5
                                                _ = v26
                                            }
                                            var v29 int8 = 127
                                            _ = v22
                                            _ = v29
                                        }
                                    }
                                case "a":
                                    switch (fi18 + (^(fi18))) {
                                    case -813:
                                        var v30 int8 = int8(v4)
                                        _ = v30
                                    default:
                                        delete(v3, (("a" + ("\n\t" + "a")) + "\n\t"))
                                    }
                                }
                            default:
                                switch -496 {
                                case -391:
                                    v9 += 389
                                case -197:
                                    var v31 int = (fi15 / (((fi15 * fi15)) | 1))
                                    fmt.Println(fi18)
                                    var v32 uint32 = 2
                                    _ = v31
                                    _ = v32
                                case 96:
                                    switch fi15 {
                                    case 127:
                                        switch fi18 {
                                        case 9223372036854775807:
                                            v5 = v14
                                        }
                                    case 0:
                                        v1 -= ((v1 / (v1 + (v1 + float64(v14)))) * 0.0)
                                    case 9223372036854775807:
                                        fmt.Println(v4, v5)
                                    default:
                                        v1 -= (3.14159 * v1)
                                    }
                                }
                            }
                            if len(v13) > 0 {
                                v13[0] = 4294967295
                            }
                        } else {
                            if (false || (v14 != ((float32(v1) / v14) / -2.25))) {
                                if v2 {
                                    var v33 uint64 = (v9 >> ((v4) & 63))
                                    if (225 > 0) {
                                        v33 = v9
                                    } else {
                                        var v34 int8 = 42
                                        _ = v34
                                    }
                                    var v35 int64 = 522
                                    _ = v33
                                    _ = v35
                                } else {
                                    var v36 int32 = 84
                                    if (0 <= 2147483647) {
                                        v37 := GenS2{F0: v2}
                                        _ = v37
                                    } else {
                                        var v38 uint64 = (v9 | (^(v9)))
                                        _ = v38
                                    }
                                    _ = v36
                                }
                                v39 := GenS4{F0: 0, F1: 255}
                                v1 = 3.14159
                                _ = v39
                            } else {
                                v40 := GenS1{F0: 951, F1: ("" + "a")}
                                v5 -= ((((v14 * float32((-2.25 + v1))) / ((v5 / v5) * (v14 / ((float32(v1) / 3.14159) * ((-2.25 + v5) * (0.0 * (v14 + (float32(((0.0 - float64(float32(v1))) + (3.14159 * (v1 / v1)))) * 0.0)))))))) + v5) - 3.14159)
                                delete(v3, "foo bar")
                                var v41 uint8 = 37
                                _ = v40
                                _ = v41
                            }
                            v42 := GenS1{F0: 2147483647, F1: (("a" + "foo bar") + ("a" + ""))}
                            _ = v42
                        }
                        v43 := GenS5{F0: fi15, F1: (850 >= 0), F2: 100}
                        if (((42 <= -1) || ((3.14159 / v1) != -2.25)) && v2) {
                            v5 = (1e10 + float32(1e10))
                            var v44 float32 = v5
                            _ = v44
                        }
                        var v45 bool = false
                        _ = v43
                        _ = v45
                    }
                    v46 := GenS1{F0: 0, F1: ("foo bar" + "foo bar")}
                    _ = v46
                }
                v14 -= (v14 * (((v14 + (1.5 / (-2.25 + v14))) * v5) - -2.25))
                if v2 {
                    var v47 int64 = 9223372036854775807
                    var v48 uint = v4
                    var v49 float32 = v5
                    _ = v47
                    _ = v48
                    _ = v49
                }
                _ = len(v3)
                _ = fi18
            }
            _ = fi15
        }
    case 716:
        for fi50 := 0; fi50 < 19; fi50++ {
            var v51 uint32 = 761
            var v52 uint8 = 1
            if v2 {
                var v53 bool = ((-1 <= -1) && (!false))
                v53 = (v51 >= (v51 | 1))
                v3["\n\t"] = fi50
                _ = v53
            }
            _ = fi50
            _ = v51
            _ = v52
        }
    case 2147483647:
        for fi54 := 0; fi54 < 8; fi54++ {
            fmt.Println(fi54, v14, v9)
            var v55 int16 = 32767
            v56 := GenS3{F0: 33, F1: (("foo bar" + "foo bar") + (("" + ("\n\t" + "foo bar")) + "")), F2: (v55 - v55)}
            _ = fi54
            _ = v55
            _ = v56
        }
    }
    for fi57 := 0; fi57 < 2; fi57++ {
        v58 := GenS1{F0: 2147483647, F1: ""}
        switch ("hello" + ((("foo bar" + "a") + "\n\t") + (("foo bar" + "a") + "hello"))) {
        case "":
            var v59 int = (fi57 & fi57)
            _ = v59
        }
        var v60 int32 = 0
        _ = fi57
        _ = v58
        _ = v60
    }
    var v61 bool = (v9 < v9)
    switch "\n\t" {
    case "foo bar":
        v4 += v4
    default:
        v62 := v3[""]
        _ = v62
    }
    v3["hello"] = -2147483648
    var v63 int16 = 1
    v63 = v63
    fmt.Println(v63, v1, v61)
    v61 = v61
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v63
    _ = v9
}
