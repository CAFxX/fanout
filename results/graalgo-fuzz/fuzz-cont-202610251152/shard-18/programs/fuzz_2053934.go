package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 int8 }
func genFn0() int32 {
    if false {
        if true {
            if true {
                var v7 uint8 = 128
                _ = v7
            } else {
                var v8 string = "a"
                if false {
                    var v9 int32 = -839
                    v9 -= -2147483648
                    for fi10 := 0; fi10 < 7; fi10++ {
                        v11 := []int{(^(fi10)), fi10, fi10}
                        _ = fi10
                        _ = v11
                    }
                    _ = v9
                } else {
                    v12 := GenS1{F0: 0, F1: ("foo bar" + v8)}
                    if false {
                        v13 := GenS1{F0: -902, F1: "\n\t"}
                        v14 := GenS2{F0: 502, F1: -1}
                        var v15 uint64 = 936
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        var v16 int64 = 0
                        var v17 uint16 = 404
                        _ = v16
                        _ = v17
                    }
                    _ = v12
                }
                v8 = "a"
                _ = v8
            }
            var v18 uint16 = 255
            _ = v18
        } else {
            v19 := []int{32767}
            if true {
                var v20 uint64 = 579
                _ = v20
            }
            for fi21 := 0; fi21 < 3; fi21++ {
                if len(v19) > 0 {
                    v19[0] = fi21
                }
                var v22 int = fi21
                for fi23 := 0; fi23 < 7; fi23++ {
                    fmt.Println(fi21, fi23)
                    _ = fi23
                }
                _ = fi21
                _ = v22
            }
            _ = v19
        }
        var v24 uint8 = 127
        _ = v24
    } else {
        v25 := GenS1{F0: 42, F1: ""}
        v26 := GenS2{F0: 4294967295, F1: 2147483647}
        v27 := []int{}
        _ = v25
        _ = v26
        _ = v27
    }
    var v28 uint16 = 886
    v29 := GenS2{F0: 580, F1: 1}
    v30 := GenS3{F0: 255, F1: 2}
    _ = v28
    _ = v29
    _ = v30
    return 0
}

func main() {
    var v1 int = -338
    var v2 uint8 = 42
    var v3 int16 = 493
    v4 := make(map[string]string)
    v5 := make(map[int]uint32)
    v6 := make(map[int]bool)
    v31 := GenS2{F0: 137, F1: 42}
    v32 := GenS2{F0: 42, F1: -789}
    switch ("hello" + ("\n\t" + "")) {
    case "\n\t":
        v1 += -900
    case "a":
        var v33 uint64 = 1
        _ = v33
    }
    fmt.Println(v1, v1, v2)
    for fi34 := 0; fi34 < 13; fi34++ {
        fmt.Println(v2)
        _ = fi34
    }
    fmt.Println(v1, v2, v3)
    v35 := GenS1{F0: -765, F1: "hello"}
    if (42 <= 4294967295) {
        var v36 int = 42
        switch (-(v36)) {
        case -9223372036854775808:
            v37 := []int{}
            _ = v37
        case -1:
            v3 += v3
        case -405:
            if false {
                for fi38 := 0; fi38 < 7; fi38++ {
                    switch -944 {
                    case -9223372036854775808:
                        fmt.Println(v1)
                    default:
                        for fi39 := 0; fi39 < 9; fi39++ {
                            var v40 float64 = 1e10
                            _ = fi39
                            _ = v40
                        }
                    }
                    for fi41 := 0; fi41 < 15; fi41++ {
                        v42 := GenS4{F0: fi41, F1: 917}
                        switch (fi38 ^ v36) {
                        case 32767:
                            var v43 float32 = -2.25
                            _ = v43
                        case -2147483648:
                            v3 += (^(v3))
                            v44 := GenS5{F0: 63, F1: v3}
                            _ = v44
                        }
                        _ = fi41
                        _ = v42
                    }
                    var v45 int32 = -2147483648
                    _ = fi38
                    _ = v45
                }
                delete(v5, (v36 & 128))
                v46 := GenS6{F0: 1}
                _ = v46
            } else {
                v47 := GenS2{F0: 2147483647, F1: -2147483648}
                _ = v47
            }
        }
        _ = v36
    }
    var v48 uint64 = 798
    v1 -= v1
    for fi49 := 0; fi49 < 20; fi49++ {
        for fi50 := 0; fi50 < 13; fi50++ {
            v2 = 128
            v1 += (fi50 & fi50)
            _ = fi50
        }
        v51 := GenS5{F0: 66, F1: (^(v3))}
        _ = fi49
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
