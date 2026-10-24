package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int8 }
type GenS2 struct { F0 uint16; F1 int8; F2 int16 }
func genFn0() uint64 {
    if (32768 <= 0) {
        switch 128 {
        case 32768:
            var v7 string = ""
            _ = v7
        }
        v8 := GenS1{F0: ("" + ("\n\t" + "hello")), F1: -55, F2: 100}
        _ = v8
    }
    v9 := GenS1{F0: "foo bar", F1: 100, F2: -47}
    var v10 float32 = 1.5
    switch -9223372036854775808 {
    case 127:
        if (!((float32((float64(v10) * 1e10)) != v10) || true)) {
            if true {
                fmt.Println(v10)
            } else {
                var v11 int32 = 1
                _ = v11
            }
            v12 := []int{399, -880, -917, 42}
            v10 = v10
            _ = v12
        }
    case 0:
        v10 = v10
    case -30:
        if ((-1 <= -32768) && (float64(0.1) <= 1e10)) {
            var v13 int32 = 1
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := []int{}
                v16 := GenS1{F0: "a", F1: 1, F2: 42}
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = v13
        } else {
            var v17 string = "a"
            v17 = (v17 + (((v17 + (("a" + (v17 + v17)) + v17)) + "a") + (((((((("\n\t" + v17) + v17) + v17) + v17) + ("a" + v17)) + v17) + ("a" + "")) + "a")))
            _ = v17
        }
    default:
        for fi18 := 0; fi18 < 16; fi18++ {
            for fi19 := 0; fi19 < 12; fi19++ {
                v20 := []int{4294967295, fi19, (fi19 >> ((723) & 63))}
                var v21 bool = false
                v22 := GenS1{F0: ("\n\t" + ("a" + "a")), F1: -60, F2: -104}
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            v23 := []int{}
            v10 = float32(1e10)
            _ = fi18
            _ = v23
        }
    }
    v24 := GenS1{F0: "foo bar", F1: 117, F2: 104}
    _ = v10
    _ = v24
    _ = v9
    return 1
}
func genFn1(p0 bool) bool {
    for fi25 := 0; fi25 < 19; fi25++ {
        p0 = (p0 && false)
        var v26 float32 = -2.25
        var v27 int8 = -128
        _ = fi25
        _ = v26
        _ = v27
    }
    if ((701 != 620) && false) {
        v28 := []int{1, 2147483647}
        _ = v28
    } else {
        if (3.14159 != ((((1.5 + (1e10 * 0.0)) * ((3.14159 + ((1e10 * 0.0) - float64(1e10))) / -2.25)) / 1.5) - (((3.14159 - 0.0) * -2.25) / 1e10))) {
            p0 = true
        }
        fmt.Println(p0)
    }
    var v29 int32 = -1
    switch ("\n\t" + "hello") {
    case "hello":
        p0 = (100 <= 100)
    case "":
        var v30 bool = false
        _ = v30
    case "\n\t":
        v29 = (v29 + v29)
    }
    _ = p0
    _ = v29
    return (1000 < -32768)
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 1
    v3 := []string{("\n\t" + "hello"), "foo bar"}
    v4 := make(map[uint32]int)
    var v5 int64 = 1
    var v6 float32 = ((1e10 - 0.1) / 0.1)
    if false {
        for fi31 := 0; fi31 < 16; fi31++ {
            v32 := v1[""]
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi31
            _ = v32
        }
        v4[65535] = 32767
        v33 := GenS1{F0: "hello", F1: -128, F2: 127}
        _ = v33
    } else {
        v5 = v5
        var v34 float32 = v6
        if true {
            v5 -= v5
            v3 = append(v3, "")
            v6 = (((v6 + v6) - 0.0) * v6)
        }
        _ = v34
    }
    var v35 uint32 = 65535
    var v36 uint16 = 2
    v35 = v35
    var v37 uint = 1
    for fi38 := 0; fi38 < 1; fi38++ {
        if (1 <= 42) {
            if true {
                var v39 uint64 = genFn0()
                _ = v39
            } else {
                v6 -= v6
            }
            fmt.Println(v35, v5, v35)
        }
        var v40 uint = 42
        v1["foo bar"] = 0
        _ = fi38
        _ = v40
    }
    var v41 bool = ((false || ((!true) && (v35 < v35))) && false)
    v42 := GenS2{F0: v36, F1: -22, F2: 1000}
    var v43 float64 = 1e10
    fmt.Println(v37, v36)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
