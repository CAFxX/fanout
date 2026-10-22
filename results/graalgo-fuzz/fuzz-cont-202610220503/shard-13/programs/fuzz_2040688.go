package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int16; F1 uint; F2 int16 }
func genFn0(p0 uint32, p1 int16) string {
    switch 2147483647 {
    case 233:
        v6 := GenS1{F0: ""}
        _ = v6
    case 9223372036854775807:
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := []int{(fi7 ^ fi7)}
            var v9 uint64 = uint64(p0)
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    case -9223372036854775808:
        v10 := GenS2{F0: 32767, F1: (^(p0))}
        _ = v10
    default:
        var v11 bool = (42 > 65535)
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            p1 += p1
            _ = fi13
        }
        v14 := []int{(^(fi12)), (-(fi12)), (^(fi12))}
        _ = fi12
        _ = v14
    }
    var v15 float64 = 0.0
    for fi16 := 0; fi16 < 19; fi16++ {
        p0 = (p0 ^ (p0 >> ((768) & 63)))
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = v15
    return "\n\t"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 0
    v3 := []string{("hello" + "\n\t"), (("foo bar" + "a") + "\n\t"), "\n\t"}
    var v4 float32 = -2.25
    v5 := make(map[uint32]int)
    if (false || false) {
        v17 := GenS1{F0: "hello"}
        _ = v17
    } else {
        if false {
            switch "hello" {
            case "foo bar":
                var v18 bool = (true && (!false))
                _ = v18
            default:
                v4 = v4
            }
            var v19 int32 = 2147483647
            var v20 bool = false
            _ = v19
            _ = v20
        }
    }
    if true {
        for fi21 := 0; fi21 < 7; fi21++ {
            switch fi21 {
            case 4294967295:
                if len(v3) > 0 {
                    v3[0] = ("\n\t" + "hello")
                }
            case 32768:
                if true {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(v4)
                }
            case 127:
                var v22 int32 = -1
                _ = v22
            default:
                v4 = ((1e10 / 1e10) - 0.1)
            }
            _ = fi21
        }
        var v23 int8 = 42
        switch ("\n\t" + "\n\t") {
        case "":
            for fi24 := 0; fi24 < 12; fi24++ {
                var v25 int = 1
                var v26 int = 42
                for fi27 := 0; fi27 < 18; fi27++ {
                    v25 += 1
                    switch fi24 {
                    case 9223372036854775807:
                        switch "" {
                        case "foo bar":
                            var v28 float64 = 3.14159
                            if true {
                                v2 = v2
                                fmt.Println(v26, fi27, v2)
                                if (299 == 2) {
                                    v29 := GenS1{F0: "\n\t"}
                                    v26 += (^(v26))
                                    v2 = int(v23)
                                    var v30 int64 = 471
                                    switch fi24 {
                                    case 907:
                                        if true {
                                            if len(v3) > 0 {
                                                v3[0] = ("a" + "")
                                            }
                                        }
                                    case 192:
                                        v31 := GenS3{F0: 56, F1: 65535, F2: 32767}
                                        _ = v31
                                    default:
                                        v23 += v23
                                    }
                                    var v32 int8 = -128
                                    var v33 int32 = 395
                                    v34 := GenS2{F0: 65535, F1: 42}
                                    var v35 bool = true
                                    _ = v29
                                    _ = v30
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                            }
                            if false {
                                v36 := v5[110]
                                switch (v25 - (v36 * 32767)) {
                                case 1:
                                    if (2147483647 <= -2147483648) {
                                        v37 := GenS3{F0: 1000, F1: 343, F2: 200}
                                        v2 = v36
                                        _ = v37
                                    }
                                case 455:
                                    var v38 string = ""
                                    _ = v38
                                default:
                                    var v39 uint16 = 2
                                    _ = v39
                                }
                                v3 = append(v3, "a")
                                _ = v36
                            }
                            _ = v28
                        case "\n\t":
                            fmt.Println(v26)
                        }
                    case -918:
                        switch (v2 & v26) {
                        case -2147483648:
                            v40 := GenS3{F0: -529, F1: 255, F2: 42}
                            _ = v40
                        case -1:
                            switch v25 {
                            case 4294967295:
                                var v41 float32 = v4
                                _ = v41
                            }
                        case 42:
                            var v42 int = (-(fi24))
                            _ = v42
                        default:
                            v23 = v23
                        }
                    }
                    _ = fi27
                }
                _ = fi24
                _ = v25
                _ = v26
            }
        case "\n\t":
            var v43 string = ("foo bar" + "")
            _ = v43
        }
        _ = v23
    }
    if (1e10 > v4) {
        if true {
            v2 -= v2
            fmt.Println(v2)
        }
        v4 = (v4 - (v4 * 1.5))
    } else {
        _ = len(v3)
        v4 += v4
    }
    v2 = v2
    var v44 int16 = 32767
    v4 += (((v4 * (-2.25 / (v4 / (((0.0 - v4) * (v4 - 3.14159)) + ((((((v4 * (v4 / ((v4 * v4) / ((v4 + v4) / -2.25)))) + v4) - v4) / 1.0) / v4) / v4))))) * float32(float64((v4 - 0.0)))) + v4)
    v45 := v1[("foo bar" + ("hello" + "hello"))]
    var v46 uint8 = 127
    if (true || (84 == 100)) {
        v45 = 2147483647
        if false {
            for fi47 := 0; fi47 < 1; fi47++ {
                fmt.Println(v45, v4)
                _ = fi47
            }
        }
    } else {
        v2 -= v2
        for fi48 := 0; fi48 < 6; fi48++ {
            var v49 uint16 = 42
            for fi50 := 0; fi50 < 2; fi50++ {
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
                _ = fi50
            }
            _ = fi48
            _ = v49
        }
        v2 += (v2 | 32768)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
