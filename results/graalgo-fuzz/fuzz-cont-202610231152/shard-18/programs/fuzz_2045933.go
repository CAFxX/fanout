package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS4 struct { F0 bool; F1 int8; F2 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 int32; F2 int }
func genFn0(p0 uint32, p1 string, p2 bool) int8 {
    var v4 float32 = (-2.25 + 3.14159)
    if p2 {
        for fi5 := 0; fi5 < 1; fi5++ {
            if (242 < 241) {
                v6 := []int{(fi5 + fi5), (fi5 & 372), fi5}
                p2 = false
                fmt.Println(v4)
                _ = v6
            } else {
                var v7 int64 = 1
                v8 := GenS1{F0: 255, F1: 255, F2: 0}
                _ = v7
                _ = v8
            }
            var v9 uint8 = 216
            v10 := []int{fi5, fi5}
            _ = fi5
            _ = v10
            _ = v9
        }
        switch ((p1 + "\n\t") + p1) {
        case "foo bar":
            v11 := []int{383}
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 100
}
func genFn1(p0 bool) string {
    if false {
        fmt.Println(p0)
        for fi12 := 0; fi12 < 8; fi12++ {
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 int64 = -793
                _ = fi13
                _ = v14
            }
            switch fi12 {
            case 9223372036854775807:
                var v15 uint64 = 353
                _ = v15
            case 892:
                fmt.Println(fi12, p0)
            case 32768:
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 int8 = genFn0(2147483648, "", false)
                    var v18 uint32 = 0
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            p0 = (!(true || ((("" + "a") + "hello") >= ("a" + "hello"))))
            _ = fi12
        }
    }
    v19 := GenS1{F0: 255, F1: 32767, F2: 65535}
    var v20 bool = (!(2 == 255))
    fmt.Println(v20, v20)
    _ = p0
    _ = v19
    _ = v20
    return "hello"
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = -833
    var v3 int = 42
    if ((v3 << ((638) & 63)) == (^(v3))) {
        var v21 uint32 = 309
        v22 := v1[2147483647]
        delete(v1, v22)
        _ = v21
        _ = v22
    } else {
        if false {
            var v23 uint = 1
            _ = v23
        } else {
            var v24 uint = 4294967295
            v2 += v2
            var v25 int64 = int64(v2)
            _ = v24
            _ = v25
        }
    }
    v26 := GenS2{F0: -9223372036854775808}
    if (!false) {
        switch -1 {
        case -2147483648:
            v27 := GenS2{F0: -9223372036854775808}
            _ = v27
        case -1:
            switch "" {
            case "a":
                v28 := []int{(^(v3)), 32768}
                _ = v28
            case "foo bar":
                switch -9223372036854775808 {
                case -224:
                    for fi29 := 0; fi29 < 17; fi29++ {
                        v30 := GenS2{F0: 592}
                        fmt.Println(fi29)
                        v3 -= (^(fi29))
                        _ = fi29
                        _ = v30
                    }
                default:
                    var v31 uint32 = 1
                    _ = v31
                }
            case "hello":
                if (42 != 42) {
                    delete(v1, 2)
                    var v32 float64 = (1.5 * (0.0 * 3.14159))
                    _ = v32
                } else {
                    var v33 uint64 = 165
                    _ = v33
                }
            }
        }
        fmt.Println(v3, v2)
        var v34 uint8 = 1
        _ = v34
    }
    fmt.Println(v3, v3)
    if (!false) {
        var v35 int64 = -708
        switch 4294967295 {
        case 951:
            v3 += v3
        }
        _ = v35
    } else {
        v36 := []int{v3, 42, v3, v3}
        var v37 uint = 4294967295
        switch "" {
        case "a":
            v38 := GenS3{F0: 207, F1: 0, F2: 42}
            _ = v38
        case "hello":
            var v39 uint64 = 345
            _ = v39
        case "":
            switch 9223372036854775807 {
            case 149:
                v40 := GenS3{F0: 1, F1: -1, F2: 128}
                _ = v40
            }
        default:
            if (true && true) {
                switch "a" {
                case "a":
                    if ((!(!true)) || true) {
                        if len(v36) >= 2 {
                            v36 = v36[1:len(v36)]
                        }
                        v3 = v3
                    } else {
                        var v41 bool = true
                        if v41 {
                            var v42 uint = (^(v37))
                            _ = v42
                        } else {
                            fmt.Println(v2, v3)
                            if (v41 || v41) {
                                for fi43 := 0; fi43 < 11; fi43++ {
                                    v44 := GenS1{F0: 65535, F1: 32767, F2: 42}
                                    if (4294967295 >= 2) {
                                        v41 = (!(423 == 2))
                                        var v45 float64 = (3.14159 + float64((1e10 + -2.25)))
                                        if ((!((v3 % (((v3 ^ int(v2))) | 1)) <= v3)) && (((1.5 == (v45 * 1.5)) || (9223372036854775807 == 0)) || (!v41))) {
                                            switch "hello" {
                                            case "foo bar":
                                                v46 := GenS4{F0: v41, F1: 42, F2: 255}
                                                _ = v46
                                            case "":
                                                switch (^(fi43)) {
                                                case -2147483648:
                                                    switch (fi43 + 2147483647) {
                                                    case 128:
                                                        var v47 string = (("hello" + "") + "hello")
                                                        _ = v47
                                                    case 32767:
                                                        v41 = true
                                                    }
                                                default:
                                                    v37 = (^(v37))
                                                }
                                            default:
                                                var v48 int32 = -2147483648
                                                _ = v48
                                            }
                                        } else {
                                            fmt.Println(v2, v41, fi43)
                                            if (((0.0 * ((((3.14159 / 0.1) + 0.1) * 1e10) * 0.0)) != 3.14159) || (42 > 1)) {
                                                if (1 > 32767) {
                                                    for fi49 := 0; fi49 < 11; fi49++ {
                                                        v50 := v1[803]
                                                        v41 = (985 > 42)
                                                        _ = fi49
                                                        _ = v50
                                                    }
                                                }
                                            } else {
                                                v51 := GenS5{F0: -89}
                                                v3 = (v3 | (fi43 | -458))
                                                _ = v51
                                            }
                                            for fi52 := 0; fi52 < 16; fi52++ {
                                                v41 = false
                                                switch ("\n\t" + ("foo bar" + ("\n\t" + (((("a" + ("a" + ("foo bar" + ("foo bar" + "\n\t")))) + "hello") + ("a" + ("foo bar" + ""))) + "a")))) {
                                                case "foo bar":
                                                    v45 -= float64((-2.25 * 1.5))
                                                case "a":
                                                    v37 = v37
                                                default:
                                                    var v53 float64 = 1.5
                                                    _ = v53
                                                }
                                                switch "\n\t" {
                                                case "\n\t":
                                                    _ = len(v1)
                                                }
                                                _ = fi52
                                            }
                                        }
                                        _ = v45
                                    } else {
                                        v3 = fi43
                                        v37 += 4294967295
                                        v54 := GenS6{F0: 42, F1: -2147483648, F2: v3}
                                        _ = v54
                                    }
                                    v2 = (v2 >> ((v37) & 63))
                                    _ = fi43
                                    _ = v44
                                }
                            } else {
                                var v55 string = ((("a" + "a") + "hello") + ("hello" + "foo bar"))
                                var v56 int = 128
                                _ = v55
                                _ = v56
                            }
                            v41 = (42 >= -128)
                        }
                        _ = v41
                    }
                }
                fmt.Println(v37, v37)
            }
        }
        _ = v36
        _ = v37
    }
    v1[913] = 2
    for fi57 := 0; fi57 < 8; fi57++ {
        if (0 != 0) {
            var v58 int16 = int16(fi57)
            _ = v58
        }
        v3 += (v3 & v3)
        _ = fi57
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
}
