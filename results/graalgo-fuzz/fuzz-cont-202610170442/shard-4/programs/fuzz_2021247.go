package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
func genFn0() int16 {
    if true {
        switch -540 {
        case -1:
            var v5 int = 0
            _ = v5
        case 128:
            if (-1 > -2147483648) {
                if (true && (32767 <= 487)) {
                    var v6 int32 = 2147483647
                    _ = v6
                } else {
                    var v7 int32 = 1
                    fmt.Println(v7)
                    _ = v7
                }
                v8 := []int{-975}
                var v9 uint64 = 42
                _ = v8
                _ = v9
            } else {
                v10 := []int{32767}
                _ = v10
            }
        case 447:
            if false {
                if false {
                    var v11 float32 = 0.1
                    for fi12 := 0; fi12 < 20; fi12++ {
                        v13 := []int{(-(fi12)), fi12}
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        _ = fi12
                        _ = v13
                    }
                    v11 += 0.0
                    _ = v11
                }
                var v14 float32 = 0.0
                _ = v14
            }
        }
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 uint64 = 2
            _ = fi15
            _ = v16
        }
    } else {
        switch "a" {
        case "\n\t":
            for fi17 := 0; fi17 < 10; fi17++ {
                for fi18 := 0; fi18 < 16; fi18++ {
                    var v19 int8 = 100
                    _ = fi18
                    _ = v19
                }
                var v20 bool = (false && (!false))
                var v21 bool = v20
                _ = fi17
                _ = v20
                _ = v21
            }
        case "a":
            var v22 int64 = 9223372036854775807
            _ = v22
        case "foo bar":
            for fi23 := 0; fi23 < 14; fi23++ {
                fmt.Println(fi23)
                var v24 uint32 = 2147483648
                for fi25 := 0; fi25 < 8; fi25++ {
                    v24 += (^(v24))
                    if true {
                        v26 := []int{}
                        _ = v26
                    } else {
                        if ((1 < 42) && false) {
                            v27 := []int{236, (fi23 - fi25), fi23, fi25}
                            _ = v27
                        } else {
                            fmt.Println(fi25, fi25, fi23)
                        }
                        v28 := GenS1{F0: 255, F1: 903}
                        fmt.Println(v24, fi23, fi25)
                        _ = v28
                    }
                    var v29 int64 = 9223372036854775807
                    var v30 string = "a"
                    v29 = 0
                    _ = fi25
                    _ = v29
                    _ = v30
                }
                _ = fi23
                _ = v24
            }
        default:
            for fi31 := 0; fi31 < 6; fi31++ {
                var v32 int32 = -630
                _ = fi31
                _ = v32
            }
        }
    }
    var v33 float32 = (-2.25 - 3.14159)
    v34 := GenS1{F0: 255, F1: 607}
    v33 = (v33 * (v33 - v33))
    v35 := []int{2147483647, -875}
    _ = v33
    _ = v34
    _ = v35
    return -688
}
func genFn1(p0 uint8) uint16 {
    p0 += p0
    switch (((((((("foo bar" + "hello") + ((("foo bar" + "hello") + "\n\t") + ("a" + "\n\t"))) + ("hello" + "")) + "hello") + ("\n\t" + ("hello" + "\n\t"))) + "a") + "a") + ("foo bar" + "\n\t")) {
    case "\n\t":
        for fi36 := 0; fi36 < 11; fi36++ {
            fmt.Println(p0, fi36)
            _ = fi36
        }
    default:
        p0 -= p0
    }
    v37 := []int{4294967295, 42, 128, 32767}
    var v38 float32 = (1.5 / ((-2.25 + (-2.25 - 3.14159)) * (3.14159 / 1.5)))
    _ = p0
    _ = v37
    _ = v38
    return 255
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    var v3 uint8 = 128
    var v4 float32 = 1.5
    v4 -= 3.14159
    v3 = (^(v3))
    var v39 uint = 593
    var v40 uint16 = 823
    var v41 int8 = 77
    var v42 uint32 = 4294967295
    var v43 string = "a"
    v4 = ((3.14159 + v4) * v4)
    v44 := []int{}
    v45 := GenS1{F0: 42, F1: (^(v40))}
    switch -490 {
    case 1:
        var v46 uint64 = 42
        _ = v46
    case -416:
        v41 = (v41 >> ((v39) & 63))
    case 32768:
        switch (v43 + v43) {
        case "\n\t":
            v41 = v41
        default:
            v43 = (((v43 + (v43 + v43)) + (v43 + v43)) + v43)
        }
    default:
        for fi47 := 0; fi47 < 11; fi47++ {
            v41 += (v41 & v41)
            _ = fi47
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
