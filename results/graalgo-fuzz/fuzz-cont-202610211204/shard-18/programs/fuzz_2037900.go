package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int16, p1 uint64) bool {
    v7 := []int{42, -1, 9223372036854775807}
    v8 := GenS1{F0: (-(p0))}
    v9 := GenS2{F0: 128}
    v10 := GenS1{F0: (p0 ^ 28)}
    for fi11 := 0; fi11 < 19; fi11++ {
        switch "\n\t" {
        case "hello":
            fmt.Println(fi11, fi11, p0)
        case "":
            var v12 string = ""
            _ = v12
        case "foo bar":
            v13 := GenS2{F0: 42}
            _ = v13
        default:
            for fi14 := 0; fi14 < 15; fi14++ {
                if len(v7) > 0 {
                    v7[0] = (fi14 ^ (fi14 - fi11))
                }
                fmt.Println(fi14, fi11)
                _ = fi14
            }
        }
        v15 := GenS2{F0: 214}
        for fi16 := 0; fi16 < 5; fi16++ {
            switch "\n\t" {
            case "\n\t":
                p0 += 42
            default:
                fmt.Println(fi11, fi11)
            }
            v7 = append(v7, (-(fi11)))
            _ = fi16
        }
        _ = fi11
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return false
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]int)
    var v3 float64 = 1e10
    var v4 bool = true
    var v5 int32 = -249
    var v6 float32 = ((-2.25 * 1e10) * float32(1.5))
    if v4 {
        v17 := v1[2]
        _ = v17
    } else {
        v18 := GenS1{F0: -964}
        if true {
            var v19 uint8 = 255
            _ = v19
        }
        v4 = (("hello" + ("" + ("" + "\n\t"))) <= ("a" + "\n\t"))
        _ = v18
    }
    switch "\n\t" {
    case "":
        if true {
            if (!v4) {
                var v20 float64 = v3
                v20 -= float64(v6)
                for fi21 := 0; fi21 < 4; fi21++ {
                    v6 += v6
                    v3 = v3
                    _ = fi21
                }
                _ = v20
            }
            v5 += v5
            var v22 float64 = (1e10 - ((v3 + 1.5) * v3))
            _ = v22
        } else {
            for fi23 := 0; fi23 < 3; fi23++ {
                switch (("" + "") + "") {
                case "":
                    v5 -= (v5 ^ (-(v5)))
                case "a":
                    fmt.Println(v4)
                }
                _ = fi23
            }
        }
    default:
        var v24 int = -301
        _ = v24
    }
    switch "" {
    case "\n\t":
        v5 += v5
    case "":
        if v4 {
            if (255 == 255) {
                v4 = (v6 > 0.1)
                var v25 uint16 = 255
                _ = v25
            }
        } else {
            var v26 uint64 = 4294967295
            _ = v26
        }
    }
    if (!((v3 >= (((1.5 + float64(v6)) - v3) / 1.5)) && v4)) {
        if v4 {
            for fi27 := 0; fi27 < 9; fi27++ {
                v6 = (0.0 - (v6 / float32(v3)))
                _ = fi27
            }
            for fi28 := 0; fi28 < 7; fi28++ {
                fmt.Println(v4)
                switch (fi28 & fi28) {
                case 128:
                    var v29 uint64 = 42
                    _ = v29
                }
                _ = fi28
            }
            v30 := GenS2{F0: 2}
            _ = v30
        } else {
            if (!((v5 ^ v5) == (v5 - v5))) {
                if v4 {
                    var v31 float32 = ((1e10 / v6) * 1e10)
                    for fi32 := 0; fi32 < 1; fi32++ {
                        if (255 != 128) {
                            for fi33 := 0; fi33 < 17; fi33++ {
                                v34 := []int{(fi32 | 9223372036854775807), (-(fi33)), fi33, fi33}
                                v35 := GenS1{F0: 0}
                                v31 = (v6 * v31)
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                            switch "" {
                            case "foo bar":
                                v36 := GenS3{F0: (v5 << ((65535) & 63))}
                                _ = v36
                            default:
                                v37 := []int{127}
                                _ = v37
                            }
                        } else {
                            fmt.Println(v5)
                        }
                        _ = fi32
                    }
                    if (v5 != v5) {
                        var v38 uint = 18446744073709551615
                        fmt.Println(v31, v31)
                        _ = v38
                    } else {
                        var v39 string = "a"
                        fmt.Println(v6)
                        switch -2147483648 {
                        case -288:
                            v40 := []int{}
                            _ = v40
                        case 128:
                            var v41 float64 = -2.25
                            _ = v41
                        case -300:
                            if ((494 != 42) || (2 != 2)) {
                                var v42 uint16 = 732
                                _ = v42
                            }
                        }
                        _ = v39
                    }
                    _ = v31
                }
                var v43 uint16 = 65535
                var v44 bool = true
                _ = v43
                _ = v44
            } else {
                v45 := GenS1{F0: 1}
                var v46 uint16 = 42
                if v4 {
                    fmt.Println(v4)
                    v47 := GenS1{F0: 42}
                    _ = v47
                }
                _ = v45
                _ = v46
            }
            var v48 uint16 = 65535
            for fi49 := 0; fi49 < 16; fi49++ {
                for fi50 := 0; fi50 < 9; fi50++ {
                    v48 += (v48 - (v48 | (v48 ^ 42)))
                    var v51 int16 = -1
                    _ = fi50
                    _ = v51
                }
                switch "foo bar" {
                case "\n\t":
                    var v52 int64 = -1
                    _ = v52
                default:
                    for fi53 := 0; fi53 < 4; fi53++ {
                        fmt.Println(v4)
                        v3 = float64(((-2.25 * 0.1) / (((1e10 - v6) * v6) / 3.14159)))
                        _ = fi53
                    }
                }
                _ = fi49
            }
            _ = v48
        }
        var v54 uint = 65535
        v6 = 0.1
        _ = v54
    }
    if true {
        var v55 int16 = -991
        if (v4 || true) {
            var v56 uint16 = 32767
            v4 = v4
            fmt.Println(v56, v5)
            _ = v56
        }
        _ = v55
    }
    for fi57 := 0; fi57 < 10; fi57++ {
        switch (^(fi57)) {
        case 9223372036854775807:
            var v58 bool = false
            _ = v58
        case 4294967295:
            v4 = false
        default:
            fmt.Println(fi57, v3)
        }
        _ = fi57
    }
    v59 := []int{32767, -1, 1, -1}
    v5 = v5
    _ = len(v2)
    if len(v59) > 0 {
        v59[0] = 32767
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v59:", len(v59), v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
