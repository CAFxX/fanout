package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 bool; F1 uint; F2 uint16 }
func genFn0() uint {
    var v4 float64 = ((1e10 / (0.1 - (1e10 + (((0.1 + (-2.25 - float64(float32(float64((1e10 - (0.1 + 1e10))))))) * (1.5 - 0.1)) + 0.1)))) / 3.14159)
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 uint16 = 2
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 53
    _ = v4
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 265
    v3 := make(map[uint32]uint32)
    switch "hello" {
    case "\n\t":
        var v8 uint32 = 4294967295
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: genFn0()}
    v10 := GenS1{F0: 42, F1: 42}
    if true {
        v2 = (-(v2))
        switch "a" {
        case "\n\t":
            v11 := GenS1{F0: 301, F1: 65535}
            _ = v11
        }
        if false {
            fmt.Println(v2)
            for fi12 := 0; fi12 < 7; fi12++ {
                v13 := GenS2{F0: 103, F1: 0}
                _ = fi12
                _ = v13
            }
            v14 := GenS1{F0: 870, F1: 0}
            _ = v14
        }
    } else {
        v2 += (-(v2))
    }
    v1["hello"] = ""
    var v15 int8 = 0
    if ((1e10 * 1.5) <= (0.0 * 1e10)) {
        var v16 int = v2
        v17 := GenS1{F0: 65535, F1: 866}
        var v18 float32 = 1e10
        _ = v16
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 10; fi19++ {
            switch "a" {
            case "\n\t":
                for fi20 := 0; fi20 < 14; fi20++ {
                    var v21 uint16 = 32767
                    var v22 uint8 = 42
                    v23 := GenS1{F0: 2147483647, F1: 33}
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
            case "":
                v15 = v15
            }
            _ = len(v1)
            v2 = -9223372036854775808
            _ = fi19
        }
        v24 := GenS2{F0: 255, F1: 42}
        _ = v24
    }
    var v25 int8 = (^(v15))
    v3[1] = 65535
    for fi26 := 0; fi26 < 1; fi26++ {
        delete(v3, 220)
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 uint8 = 127
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 5; fi29++ {
            fmt.Println(fi26, fi26)
            if (("\n\t" + "foo bar") >= ("\n\t" + "\n\t")) {
                v30 := []int{}
                var v31 float32 = ((0.1 - 3.14159) + 1e10)
                v32 := GenS3{F0: 42, F1: false}
                _ = v30
                _ = v31
                _ = v32
            } else {
                v33 := v3[1]
                switch fi29 {
                case 2147483647:
                    var v34 int8 = v15
                    _ = v34
                case 365:
                    switch "a" {
                    case "hello":
                        v35 := []int{(fi26 & 128)}
                        _ = v35
                    case "\n\t":
                        if (!("" < "hello")) {
                            var v36 uint32 = v33
                            _ = v36
                        } else {
                            var v37 int32 = -1
                            var v38 int16 = 1
                            _ = v37
                            _ = v38
                        }
                    case "foo bar":
                        v25 -= 42
                    default:
                        v15 = (v15 - -37)
                    }
                }
                _ = v33
            }
            fmt.Println(v2, v25, fi26)
            var v39 string = "hello"
            _ = fi29
            _ = v39
        }
        _ = fi26
    }
    if true {
        v40 := GenS3{F0: 0, F1: false}
        v25 = v15
        var v41 int = v2
        _ = v40
        _ = v41
    } else {
        v42 := GenS2{F0: 1, F1: 42}
        _ = v42
    }
    v25 = -128
    var v43 bool = false
    var v44 uint32 = 386
    if v43 {
        if ((v43 || true) || (0 < 14)) {
            switch (("\n\t" + "\n\t") + "foo bar") {
            case "":
                var v45 uint32 = v44
                _ = v45
            case "foo bar":
                var v46 uint32 = 936
                _ = v46
            case "\n\t":
                v43 = true
            }
            _ = len(v3)
        } else {
            if (4294967295 != 14) {
                var v47 bool = (2 > 42)
                v48 := []int{v2, v2, v2, (^(v2))}
                v44 = (v44 - (v44 % ((v44) | 1)))
                _ = v47
                _ = v48
            }
            v2 -= v2
            fmt.Println(v15, v25, v2)
        }
        if (true && v43) {
            if v43 {
                delete(v1, "\n\t")
                var v49 uint16 = 2
                _ = v49
            } else {
                fmt.Println(v15, v25, v43)
            }
            var v50 string = ""
            _ = v50
        } else {
            v51 := GenS3{F0: 42, F1: v43}
            _ = v51
        }
    } else {
        v52 := GenS4{F0: (v43 && false), F1: genFn0(), F2: 32767}
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v43
    _ = v44
    _ = v9
}
