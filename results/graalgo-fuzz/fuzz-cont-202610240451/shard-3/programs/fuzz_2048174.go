package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int; F1 bool; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint; F2 int }
type GenS4 struct { F0 bool; F1 uint32; F2 uint16 }
func genFn0() uint16 {
    var v5 int8 = 1
    var v6 uint64 = 509
    if false {
        for fi7 := 0; fi7 < 12; fi7++ {
            switch "\n\t" {
            case "hello":
                var v8 float64 = 1e10
                _ = v8
            default:
                var v9 int64 = 565
                _ = v9
            }
            v5 = v5
            v10 := []int{2147483647, 4294967295}
            _ = fi7
            _ = v10
        }
    } else {
        switch ("\n\t" + (("hello" + "\n\t") + "\n\t")) {
        case "a":
            var v11 string = (("hello" + "hello") + "a")
            _ = v11
        case "":
            switch "foo bar" {
            case "":
                for fi12 := 0; fi12 < 7; fi12++ {
                    switch ("foo bar" + ("a" + "a")) {
                    case "a":
                        v13 := []int{fi12, fi12, (^(fi12)), 591}
                        _ = v13
                    default:
                        var v14 int8 = -24
                        _ = v14
                    }
                    _ = fi12
                }
            }
        case "hello":
            fmt.Println(v5)
        }
        switch 32767 {
        case 4294967295:
            switch -158 {
            case -9223372036854775808:
                v15 := GenS1{F0: "\n\t", F1: 4294967295, F2: 255}
                _ = v15
            case 42:
                v16 := []int{-798}
                _ = v16
            case 9223372036854775807:
                var v17 string = "a"
                _ = v17
            default:
                if true {
                    var v18 int64 = -9223372036854775808
                    var v19 float32 = ((((0.1 / 1e10) - float32(((0.0 - 1e10) - 0.1))) + (0.1 - (1.5 / 1.5))) + 0.1)
                    v20 := GenS1{F0: "a", F1: (v6 * 18446744073709551615), F2: 65535}
                    _ = v18
                    _ = v19
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 14; fi21++ {
                        v22 := GenS2{F0: (^(fi21)), F1: (643 >= v6), F2: 511}
                        _ = fi21
                        _ = v22
                    }
                    v23 := []int{}
                    _ = v23
                }
            }
        case -221:
            switch "foo bar" {
            case "":
                if (42 < 128) {
                    v5 -= (v5 & v5)
                    switch "hello" {
                    case "a":
                        v5 = (v5 + (^(v5)))
                    case "hello":
                        for fi24 := 0; fi24 < 10; fi24++ {
                            v6 = v6
                            _ = fi24
                        }
                    default:
                        v25 := GenS3{F0: ((1.5 - (1e10 / 3.14159)) >= -2.25), F1: 472, F2: 32767}
                        _ = v25
                    }
                } else {
                    var v26 int64 = -945
                    switch ("\n\t" + "hello") {
                    case "foo bar":
                        v5 = v5
                    case "a":
                        var v27 int = -1
                        _ = v27
                    }
                    _ = v26
                }
            default:
                if (1 > 1) {
                    v28 := GenS3{F0: (1 >= 716), F1: 255, F2: 0}
                    _ = v28
                } else {
                    var v29 uint64 = 0
                    _ = v29
                }
            }
        }
    }
    _ = v5
    _ = v6
    return 32767
}

func main() {
    var v1 int64 = 42
    var v2 bool = false
    v3 := make(map[uint32]bool)
    var v4 int8 = -15
    v1 = (v1 + v1)
    var v30 int8 = -123
    v4 = v30
    for fi31 := 0; fi31 < 18; fi31++ {
        v32 := []int{}
        v2 = false
        _ = fi31
        _ = v32
    }
    v33 := GenS3{F0: v2, F1: 0, F2: -979}
    v34 := GenS1{F0: "hello", F1: 4294967295, F2: 42}
    var v35 int16 = 1000
    var v36 int = 32768
    var v37 int = v36
    fmt.Println(v36, v2, v1)
    delete(v3, 257)
    if false {
        v1 = (v1 & v1)
        v38 := GenS3{F0: (false && (((3.14159 / 1.0) - float64(-2.25)) >= (-2.25 / 3.14159))), F1: 65535, F2: 596}
        for fi39 := 0; fi39 < 14; fi39++ {
            switch fi39 {
            case 127:
                v35 = int16(v30)
            case -118:
                v40 := []int{42, (v36 / ((32768) | 1)), (v36 + (v36 - (fi39 % (((-(fi39))) | 1)))), v36}
                _ = v40
            }
            v41 := GenS4{F0: (true || v2), F1: 2, F2: 218}
            _ = fi39
            _ = v41
        }
        _ = v38
    }
    v42 := GenS4{F0: false, F1: 0, F2: 32767}
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v42
}
