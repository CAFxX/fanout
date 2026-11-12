package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int8 }
func genFn0() int32 {
    if true {
        var v4 uint8 = 42
        _ = v4
    } else {
        for fi5 := 0; fi5 < 1; fi5++ {
            if ("" == "hello") {
                var v6 int32 = 0
                var v7 uint64 = 4294967295
                _ = v6
                _ = v7
            } else {
                if true {
                    switch "a" {
                    case "foo bar":
                        var v8 int32 = 938
                        _ = v8
                    case "\n\t":
                        var v9 uint8 = 57
                        _ = v9
                    case "hello":
                        fmt.Println(fi5)
                    default:
                        var v10 int = 9223372036854775807
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 9; fi11++ {
                        var v12 uint32 = 4294967295
                        _ = fi11
                        _ = v12
                    }
                    if (42 <= 2) {
                        var v13 uint32 = 531
                        _ = v13
                    }
                } else {
                    v14 := GenS1{F0: ((((("" + "") + "\n\t") + "a") + "") + (("a" + "a") + "\n\t")), F1: 65535, F2: 129}
                    _ = v14
                }
            }
            v15 := []int{(fi5 >> ((681) & 63)), (fi5 + fi5)}
            _ = fi5
            _ = v15
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 string = "foo bar"
            for fi18 := 0; fi18 < 16; fi18++ {
                v17 = v17
                var v19 uint32 = 42
                v20 := []int{(^(fi18)), fi16, (^(fi18))}
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi16
            _ = v17
        }
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        for fi22 := 0; fi22 < 4; fi22++ {
            if (1 < 42) {
                v23 := GenS2{F0: -50}
                var v24 float32 = (-2.25 / 1e10)
                v24 = v24
                _ = v23
                _ = v24
            }
            _ = fi22
        }
        v25 := []int{}
        _ = fi21
        _ = v25
    }
    var v26 string = ("foo bar" + "a")
    for fi27 := 0; fi27 < 1; fi27++ {
        v28 := []int{fi27, (^(fi27)), fi27}
        v29 := GenS1{F0: (v26 + ("hello" + "hello")), F1: 2147483647, F2: 48}
        fmt.Println(v26, v26)
        _ = fi27
        _ = v28
        _ = v29
    }
    _ = v26
    return -500
}

func main() {
    var v1 bool = true
    var v2 uint = 958
    v3 := []int{2147483647}
    for fi30 := 0; fi30 < 6; fi30++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v31 uint16 = 32767
        _ = fi30
        _ = v31
    }
    var v32 uint64 = 42
    if (-32768 <= -32768) {
        v33 := GenS2{F0: 100}
        var v34 uint32 = 1
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        v1 = (v32 == 42)
        _ = fi35
    }
    var v36 uint = 42
    if false {
        v3 = append(v3, 2147483647)
        var v37 int64 = 9223372036854775807
        _ = v37
    }
    v3 = append(v3, 4294967295)
    switch 9223372036854775807 {
    case 32767:
        v36 = (v36 << ((v2) & 63))
    case -1:
        v36 -= (v36 ^ 42)
    }
    var v38 uint8 = 255
    fmt.Println(v1, v32, v38)
    switch 127 {
    case -2147483648:
        var v39 uint64 = v32
        _ = v39
    }
    var v40 uint8 = 2
    v1 = ((float64(float32(0.0)) > 1.5) && v1)
    switch 0 {
    case -67:
        if false {
            if len(v3) > 0 {
                v3[0] = 127
            }
            switch 42 {
            case 9223372036854775807:
                fmt.Println(v1, v2)
            }
            var v41 int8 = 127
            _ = v41
        } else {
            fmt.Println(v40, v38)
        }
    case -18:
        var v42 int16 = 32767
        _ = v42
    }
    switch (("foo bar" + ("" + (("" + "") + ""))) + "hello") {
    case "a":
        var v43 uint16 = 410
        _ = v43
    case "\n\t":
        if ((!v1) && v1) {
            v36 = (^(v36))
            if len(v3) > 0 {
                v3[0] = -754
            }
            var v44 bool = v1
            _ = v44
        } else {
            v32 += v32
            switch "foo bar" {
            case "":
                for fi45 := 0; fi45 < 14; fi45++ {
                    v38 -= v38
                    v38 = v40
                    _ = fi45
                }
            case "a":
                var v46 string = "a"
                _ = v46
            case "hello":
                var v47 int32 = 186
                _ = v47
            default:
                var v48 uint16 = 32767
                _ = v48
            }
            for fi49 := 0; fi49 < 18; fi49++ {
                v50 := GenS2{F0: 127}
                v51 := GenS2{F0: 85}
                _ = fi49
                _ = v50
                _ = v51
            }
        }
    }
    switch 1 {
    case -2147483648:
        fmt.Println(v38, v36)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v38
    _ = v40
}
