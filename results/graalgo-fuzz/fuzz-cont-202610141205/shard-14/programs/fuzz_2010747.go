package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 bool; F1 int; F2 uint32 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    var v3 int16 = -1
    var v4 string = ""
    v5 := make(map[string]int)
    switch "a" {
    case "":
        var v6 uint64 = 2
        _ = v6
    case "a":
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int64 = 0
            _ = fi7
            _ = v8
        }
    case "\n\t":
        if (32768 < 0) {
            var v9 float32 = 3.14159
            _ = v9
        } else {
            v10 := GenS1{F0: (((("foo bar" + v4) + v4) + "hello") + v4), F1: 1}
            if (255 <= 0) {
                switch 42 {
                case -1:
                    v11 := GenS1{F0: v4, F1: 328}
                    _ = v11
                case 4294967295:
                    for fi12 := 0; fi12 < 1; fi12++ {
                        v13 := []int{}
                        _ = fi12
                        _ = v13
                    }
                case 395:
                    for fi14 := 0; fi14 < 17; fi14++ {
                        v3 += int16(fi14)
                        _ = fi14
                    }
                default:
                    v3 = (v3 + (v3 + (^(v3))))
                }
                v15 := GenS1{F0: v4, F1: 4294967295}
                _ = v15
            } else {
                fmt.Println(v3, v4)
                for fi16 := 0; fi16 < 6; fi16++ {
                    switch fi16 {
                    case 4294967295:
                        var v17 bool = true
                        _ = v17
                    case -9223372036854775808:
                        v18 := []int{}
                        _ = v18
                    default:
                        v19 := GenS1{F0: ("foo bar" + v4), F1: 4294967295}
                        _ = v19
                    }
                    var v20 uint = 4294967295
                    v21 := []int{fi16}
                    _ = fi16
                    _ = v20
                    _ = v21
                }
            }
            switch 790 {
            case 9223372036854775807:
                fmt.Println(v4, v3)
            default:
                if (!(0 >= -1)) {
                    v22 := GenS1{F0: "", F1: 0}
                    if false {
                        if (false || true) {
                            var v23 float64 = 0.1
                            fmt.Println(v4)
                            v23 -= v23
                            _ = v23
                        } else {
                            v24 := []int{42, 0}
                            if ((((((("\n\t" + "hello") + (v4 + (("a" + "\n\t") + ""))) + ("hello" + v4)) + v4) + (("a" + (("a" + v4) + (v4 + ""))) + v4)) + "a") != v4) {
                                var v25 int = 127
                                fmt.Println(v4, v25)
                                _ = v25
                            }
                            v4 = ("\n\t" + v4)
                            _ = v24
                        }
                    }
                    _ = v22
                } else {
                    for fi26 := 0; fi26 < 7; fi26++ {
                        v27 := GenS1{F0: "", F1: 2}
                        _ = len(v1)
                        v28 := v2["foo bar"]
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    v29 := v1["foo bar"]
                    var v30 float32 = 1e10
                    _ = v29
                    _ = v30
                }
            }
            _ = v10
        }
    }
    v3 = v3
    fmt.Println(v3, v4)
    if (1 <= 825) {
        for fi31 := 0; fi31 < 1; fi31++ {
            var v32 uint8 = 2
            switch ((v4 + v4) + v4) {
            case "foo bar":
                v32 = v32
            }
            v4 = v4
            _ = fi31
            _ = v32
        }
    } else {
        _ = len(v5)
    }
    var v33 float64 = 1e10
    for fi34 := 0; fi34 < 17; fi34++ {
        for fi35 := 0; fi35 < 10; fi35++ {
            if (128 != 28) {
                switch (fi34 + (fi34 >> ((4294967295) & 63))) {
                case 355:
                    v4 = v4
                case -2147483648:
                    switch v4 {
                    case "":
                        if (true || (false || (float32(v33) >= 0.1))) {
                            v36 := v1[v4]
                            _ = v36
                        } else {
                            v3 -= (v3 >> ((787) & 63))
                            v37 := []int{(-(fi34)), 0}
                            var v38 int32 = 1
                            _ = v37
                            _ = v38
                        }
                    case "hello":
                        if true {
                            var v39 int8 = 100
                            v3 = (-(v3))
                            v4 = ("foo bar" + v4)
                            _ = v39
                        }
                    case "\n\t":
                        v4 = v4
                    default:
                        var v40 uint32 = 2147483647
                        _ = v40
                    }
                    var v41 float32 = 3.14159
                    if true {
                        if false {
                            switch fi35 {
                            case -2147483648:
                                var v42 uint16 = 42
                                _ = v42
                            case 128:
                                var v43 uint = 42
                                delete(v5, v4)
                                var v44 string = ("hello" + (v4 + v4))
                                _ = v43
                                _ = v44
                            case 4294967295:
                                v45 := v1[v4]
                                _ = v45
                            }
                        } else {
                            var v46 bool = true
                            _ = v46
                        }
                        v47 := GenS1{F0: "a", F1: 1}
                        var v48 int16 = (v3 & (-(v3)))
                        _ = v47
                        _ = v48
                    } else {
                        if (((513 > 2147483647) || false) && true) {
                            v1[v4] = fi34
                        } else {
                            v33 = float64((0.0 / -2.25))
                        }
                        var v49 uint8 = 255
                        _ = v49
                    }
                    _ = v41
                }
                if false {
                    fmt.Println(fi34, v3)
                    v3 += 42
                    _ = len(v5)
                } else {
                    delete(v1, (v4 + v4))
                    switch ((v4 + "hello") + (v4 + (v4 + "hello"))) {
                    case "":
                        v50 := GenS1{F0: "a", F1: 1}
                        _ = v50
                    case "foo bar":
                        v3 -= 32767
                    }
                }
                var v51 uint32 = 65535
                v1[v4] = 4294967295
                var v52 string = v4
                _ = v51
                _ = v52
            } else {
                if (259 > 32767) {
                    v53 := GenS1{F0: "", F1: 1}
                    _ = v53
                } else {
                    v54 := GenS2{F0: 2, F1: -9223372036854775808}
                    _ = v54
                }
                if ((127 != 127) || (!(v33 > (v33 + ((0.0 + (v33 * (((3.14159 - ((float64(0.1) - v33) + (((-2.25 / v33) / v33) * 1.5))) / v33) * v33))) - v33))))) {
                    switch fi34 {
                    case 4294967295:
                        _ = len(v1)
                    case -9223372036854775808:
                        var v55 bool = (true || true)
                        _ = v55
                    default:
                        fmt.Println(fi35, v4, fi35)
                    }
                } else {
                    var v56 float32 = -2.25
                    _ = v56
                }
            }
            switch fi35 {
            case 4294967295:
                switch fi34 {
                case -983:
                    v2[(("" + (v4 + "hello")) + v4)] = v4
                }
            case 42:
                v57 := GenS3{F0: true, F1: (fi34 + 4294967295), F2: 4294967295}
                _ = v57
            case -1:
                v33 += (0.0 + 0.0)
            }
            v4 = v4
            var v58 string = ("\n\t" + "")
            _ = fi35
            _ = v58
        }
        for fi59 := 0; fi59 < 19; fi59++ {
            fmt.Println(v4, fi34, fi34)
            switch fi34 {
            case -437:
                v4 = ""
            }
            var v60 uint64 = 0
            _ = fi59
            _ = v60
        }
        _ = fi34
    }
    for fi61 := 0; fi61 < 16; fi61++ {
        v33 += -2.25
        if true {
            v62 := GenS3{F0: false, F1: (fi61 + (fi61 + (fi61 << ((42) & 63)))), F2: 622}
            _ = v62
        } else {
            if true {
                v63 := GenS3{F0: false, F1: 32768, F2: 482}
                for fi64 := 0; fi64 < 11; fi64++ {
                    v65 := []int{fi61, fi61, fi64}
                    _ = fi64
                    _ = v65
                }
                _ = v63
            } else {
                v66 := []int{(fi61 * fi61), fi61, fi61}
                v67 := GenS1{F0: "foo bar", F1: 4294967295}
                delete(v2, ("a" + (v4 + v4)))
                _ = v66
                _ = v67
            }
        }
        var v68 int32 = -2147483648
        _ = fi61
        _ = v68
    }
    v69 := GenS2{F0: 65535, F1: -9223372036854775808}
    for fi70 := 0; fi70 < 6; fi70++ {
        var v71 uint = 255
        _ = fi70
        _ = v71
    }
    var v72 float32 = float32((3.14159 / float64(float32(1e10))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v69
    _ = v72
}
