package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string }
func genFn0(p0 bool, p1 int64, p2 int8) bool {
    fmt.Println(p1, p2)
    switch "" {
    case "":
        p1 = 997
    case "foo bar":
        for fi5 := 0; fi5 < 8; fi5++ {
            switch fi5 {
            case 0:
                for fi6 := 0; fi6 < 20; fi6++ {
                    if p0 {
                        var v7 uint = 42
                        _ = v7
                    }
                    _ = fi6
                }
            default:
                v8 := GenS1{F0: -1}
                _ = v8
            }
            v9 := []int{-193, -1, (fi5 | (fi5 + -52)), fi5}
            _ = fi5
            _ = v9
        }
    case "a":
        var v10 int16 = 1
        _ = v10
    default:
        var v11 string = ("hello" + ("" + "a"))
        _ = v11
    }
    p0 = false
    _ = p0
    _ = p1
    _ = p2
    return false
}
func genFn1(p0 uint16, p1 uint) uint16 {
    switch 32767 {
    case 32768:
        v12 := []int{32768, 127, -477, -9223372036854775808}
        _ = v12
    case 4294967295:
        v13 := GenS1{F0: 1}
        _ = v13
    default:
        var v14 uint8 = 1
        _ = v14
    }
    switch 804 {
    case 0:
        var v15 uint8 = 2
        _ = v15
    case 1:
        if (18446744073709551615 < 4294967295) {
            switch "hello" {
            case "foo bar":
                if (42 == 121) {
                    if true {
                        v16 := []int{4294967295, 1, -9223372036854775808}
                        _ = v16
                    } else {
                        p1 -= (p1 << ((p1) & 63))
                        fmt.Println(p0, p1)
                        v17 := GenS1{F0: 42}
                        _ = v17
                    }
                    v18 := GenS1{F0: 326}
                    _ = v18
                }
            case "":
                if (628 >= 0) {
                    v19 := GenS2{F0: 927}
                    v20 := []int{1, -477, 32768}
                    switch 42 {
                    case 0:
                        switch 32767 {
                        case -2147483648:
                            v20 = append(v20, -2147483648)
                        case -794:
                            if false {
                                var v21 uint32 = 65535
                                _ = v21
                            }
                        case -209:
                            if (true && (((((3.14159 / 3.14159) * (3.14159 * 3.14159)) + 1e10) >= 0.0) && false)) {
                                var v22 float64 = 1e10
                                v23 := GenS1{F0: -1}
                                for fi24 := 0; fi24 < 10; fi24++ {
                                    _ = len(v20)
                                    v22 = (v22 + v22)
                                    _ = fi24
                                }
                                _ = v22
                                _ = v23
                            }
                        }
                    case 42:
                        var v25 int = -334
                        _ = v25
                    }
                    _ = v19
                    _ = v20
                }
            }
            if (false || ((("hello" + ("\n\t" + "")) <= ("foo bar" + ("foo bar" + ("a" + (("" + "a") + "hello"))))) && (-816 != 1))) {
                p1 = (p1 & p1)
                if true {
                    for fi26 := 0; fi26 < 16; fi26++ {
                        v27 := []int{}
                        _ = fi26
                        _ = v27
                    }
                    p1 -= p1
                    v28 := []int{128}
                    _ = v28
                }
            }
        }
    case -615:
        var v29 int16 = 32767
        _ = v29
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 628
    v4 := make(map[string]int)
    fmt.Println(v1, v1)
    if (1 == -474) {
        if (0 != 686) {
            if (!((!true) && (33 == 0))) {
                v30 := []int{660, 1, -9223372036854775808}
                _ = v30
            }
        } else {
            v1 = "foo bar"
        }
        v3 = (v3 * 616)
        for fi31 := 0; fi31 < 8; fi31++ {
            v3 += 42
            _ = fi31
        }
    } else {
        switch v1 {
        case "hello":
            if (42 > 94) {
                var v32 int = 9223372036854775807
                if (148 <= 0) {
                    v3 = 0
                } else {
                    var v33 int64 = 904
                    for fi34 := 0; fi34 < 6; fi34++ {
                        var v35 uint64 = 42
                        var v36 uint8 = 255
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    var v37 bool = true
                    _ = v33
                    _ = v37
                }
                var v38 int64 = 0
                _ = v32
                _ = v38
            }
        case "":
            fmt.Println(v3)
        }
        v39 := GenS2{F0: -251}
        var v40 uint64 = 0
        _ = v39
        _ = v40
    }
    fmt.Println(v1, v3)
    v3 += 940
    for fi41 := 0; fi41 < 20; fi41++ {
        v42 := GenS3{F0: v1}
        _ = fi41
        _ = v42
    }
    v3 = 32767
    fmt.Println(v3, v3)
    v3 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
