package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int }
func genFn0() uint8 {
    var v6 string = "hello"
    v6 = v6
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint32 = 4294967295
        var v9 int16 = 754
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        fmt.Println(v6)
        _ = fi10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        switch v6 {
        case "":
            if ((3.14159 + 1.5) <= 1.5) {
                if (!false) {
                    var v12 bool = (32767 <= 0)
                    _ = v12
                } else {
                    v13 := GenS1{F0: 42, F1: 25, F2: 1}
                    if true {
                        for fi14 := 0; fi14 < 13; fi14++ {
                            switch v6 {
                            case "a":
                                if (0.1 != (1e10 * 1.5)) {
                                    if false {
                                        v6 = "\n\t"
                                        var v15 uint32 = 949
                                        _ = v15
                                    }
                                    v16 := GenS1{F0: -684, F1: 32767, F2: -1}
                                    var v17 string = "a"
                                    _ = v16
                                    _ = v17
                                } else {
                                    switch "foo bar" {
                                    case "foo bar":
                                        if false {
                                            fmt.Println(fi14, fi11, fi14)
                                            v6 = ((v6 + v6) + v6)
                                        }
                                    case "":
                                        if ((18446744073709551615 != 18446744073709551615) && (42 <= 65535)) {
                                            var v18 int = (fi14 ^ 127)
                                            v19 := []int{1, 1, -882}
                                            _ = v18
                                            _ = v19
                                        }
                                    }
                                    fmt.Println(fi14, fi14, fi14)
                                }
                            case "\n\t":
                                v6 = v6
                            case "foo bar":
                                var v20 int32 = 42
                                _ = v20
                            }
                            var v21 uint64 = 4294967295
                            v22 := GenS2{F0: 2}
                            _ = fi14
                            _ = v21
                            _ = v22
                        }
                        var v23 float64 = 1.5
                        _ = v23
                    }
                    v24 := GenS3{F0: 255}
                    _ = v13
                    _ = v24
                }
            } else {
                for fi25 := 0; fi25 < 16; fi25++ {
                    if (true && ((3.14159 / 0.1) <= 1.5)) {
                        v26 := GenS1{F0: -1, F1: 42, F2: -697}
                        v6 = ("hello" + ("foo bar" + ((v6 + v6) + (v6 + (v6 + "a")))))
                        v6 = v6
                        if false {
                            switch v6 {
                            case "\n\t":
                                v6 = v6
                                switch fi11 {
                                case -1:
                                    switch "hello" {
                                    case "a":
                                        fmt.Println(fi25, fi25)
                                    default:
                                        v6 = v6
                                    }
                                }
                                v6 = ("" + ("" + (v6 + (v6 + "foo bar"))))
                            }
                            v6 = (v6 + v6)
                            if (-1 == 9223372036854775807) {
                                v6 = v6
                                v27 := []int{(fi25 >> ((255) & 63)), (fi25 | (fi25 * (fi25 / (((fi11 % ((fi11) | 1))) | 1)))), (fi11 << ((18446744073709551615) & 63))}
                                v6 = v6
                                _ = v27
                            }
                        } else {
                            switch "a" {
                            case "hello":
                                v28 := GenS2{F0: 42}
                                _ = v28
                            }
                            v6 = (v6 + ("foo bar" + "hello"))
                            switch v6 {
                            case "a":
                                v29 := []int{fi25, (fi25 << ((216) & 63)), (fi11 - (fi25 - 4294967295))}
                                _ = v29
                            }
                            var v30 uint16 = 1
                            _ = v30
                        }
                        _ = v26
                    }
                    _ = fi25
                }
            }
        case "\n\t":
            fmt.Println(fi11)
        }
        v31 := []int{}
        _ = fi11
        _ = v31
    }
    _ = v6
    return 42
}
func genFn1(p0 string) int8 {
    v32 := GenS2{F0: 382}
    var v33 int = -2147483648
    p0 = p0
    var v34 int16 = 1000
    var v35 int = v33
    _ = p0
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    return 0
}

func main() {
    var v1 uint16 = 0
    var v2 uint64 = 2
    v3 := []int{-1, int(v1), 2147483647, 113}
    var v4 int = 4294967295
    var v5 uint64 = (v2 ^ v2)
    var v36 uint = 4294967295
    var v37 float64 = 1e10
    if ((("hello" + ("\n\t" + "")) + "hello") >= ("hello" + (("\n\t" + ("\n\t" + "a")) + ("hello" + "hello")))) {
        v4 = v4
    }
    if ((100 < -1) || true) {
        fmt.Println(v4)
        if true {
            for fi38 := 0; fi38 < 17; fi38++ {
                v39 := GenS3{F0: v1}
                v1 = (v1 ^ 42)
                var v40 int32 = 844
                _ = fi38
                _ = v39
                _ = v40
            }
            var v41 float64 = (1.5 - (v37 + v37))
            _ = v41
        }
    }
    var v42 int8 = 88
    v43 := GenS4{F0: 400, F1: v4}
    fmt.Println(v42, v42, v37)
    v1 = (v1 * 65535)
    var v44 int = (^(v4))
    if false {
        v45 := GenS2{F0: v36}
        _ = v45
    } else {
        v42 = v42
    }
    v1 += v1
    v46 := GenS3{F0: (^(v1))}
    if (!true) {
        switch (^(v4)) {
        case 127:
            for fi47 := 0; fi47 < 9; fi47++ {
                v48 := GenS1{F0: -1, F1: (v1 - 0), F2: 9223372036854775807}
                v1 = v1
                _ = fi47
                _ = v48
            }
        case 128:
            v42 += genFn1("\n\t")
        default:
            var v49 bool = false
            _ = v49
        }
    }
    var v50 string = ("hello" + ((((("\n\t" + ("hello" + "foo bar")) + "") + "foo bar") + "") + ""))
    v51 := GenS3{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v46
    _ = v5
    _ = v50
    _ = v51
}
