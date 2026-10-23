package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int16; F1 int; F2 int16 }
type GenS3 struct { F0 int; F1 int32; F2 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64 }
func genFn0() uint32 {
    var v5 int = 0
    if true {
        var v6 string = ("hello" + "a")
        _ = v6
    } else {
        var v7 uint16 = 32767
        _ = v7
    }
    var v8 uint16 = 609
    var v9 bool = false
    var v10 uint64 = 4294967295
    _ = v10
    _ = v5
    _ = v8
    _ = v9
    return 1
}
func genFn1() uint64 {
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 uint = 931
        _ = fi11
        _ = v12
    }
    switch 128 {
    case 1:
        for fi13 := 0; fi13 < 15; fi13++ {
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 uint64 = 219
                _ = fi14
                _ = v15
            }
            v16 := GenS1{F0: -89, F1: genFn0()}
            _ = fi13
            _ = v16
        }
    case 0:
        switch 4294967295 {
        case -356:
            var v17 bool = (891 >= 2)
            _ = v17
        case 127:
            var v18 uint64 = 0
            _ = v18
        case 128:
            var v19 int = 42
            _ = v19
        }
    }
    var v20 int = -1
    switch v20 {
    case -1:
        v20 = (v20 >> ((879) & 63))
    case 0:
        v21 := GenS1{F0: -125, F1: 2147483647}
        _ = v21
    case 701:
        v20 = -2147483648
    }
    if (((3.14159 * 1.5) >= ((((0.0 - ((0.0 * (3.14159 / (1.5 / 0.1))) + 0.1)) * -2.25) / float32(3.14159)) * (1e10 - (-2.25 / 1.0)))) && true) {
        v22 := []int{}
        var v23 float32 = float32(1.5)
        _ = v22
        _ = v23
    }
    _ = v20
    return 42
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]bool)
    var v3 uint16 = 0
    var v4 int64 = 584
    v24 := GenS2{F0: -657, F1: 127, F2: 32767}
    fmt.Println(v3)
    var v25 string = ("\n\t" + "foo bar")
    if false {
        for fi26 := 0; fi26 < 15; fi26++ {
            v27 := []int{2147483647, fi26, fi26, 32768}
            v4 = v4
            _ = len(v2)
            _ = fi26
            _ = v27
        }
    } else {
        var v28 float32 = 3.14159
        if false {
            if false {
                if false {
                    v29 := []int{32768}
                    var v30 uint = 916
                    if (!false) {
                        v31 := GenS2{F0: 56, F1: 1, F2: 741}
                        if (!(!false)) {
                            v25 = ""
                            v32 := GenS3{F0: 65, F1: 2147483647, F2: 42}
                            _ = v32
                        } else {
                            v25 = v25
                            v33 := v2[4294967295]
                            v34 := GenS1{F0: 1, F1: 1}
                            _ = v33
                            _ = v34
                        }
                        if len(v29) > 0 {
                            v29[0] = 42
                        }
                        _ = v31
                    }
                    _ = v29
                    _ = v30
                }
                _ = len(v2)
            }
            var v35 int16 = -32768
            _ = v35
        }
        _ = v28
    }
    var v36 int64 = v4
    if (true && (-128 > 1)) {
        for fi37 := 0; fi37 < 2; fi37++ {
            v38 := []int{fi37, fi37, 108}
            fmt.Println(v1, v4, v1)
            v39 := GenS3{F0: 0, F1: 720, F2: 1}
            _ = fi37
            _ = v38
            _ = v39
        }
        fmt.Println(v36, v3)
        if (!true) {
            var v40 uint = 1
            _ = v40
        }
    } else {
        var v41 int64 = 0
        _ = v41
    }
    var v42 float64 = 0.0
    if true {
        var v43 string = v25
        _ = v43
    } else {
        if (876 != 165) {
            v44 := GenS2{F0: 230, F1: 128, F2: -1}
            _ = v44
        }
        v36 += (v4 | v4)
    }
    if (178 > 32768) {
        for fi45 := 0; fi45 < 18; fi45++ {
            for fi46 := 0; fi46 < 20; fi46++ {
                var v47 float64 = 0.1
                _ = fi46
                _ = v47
            }
            _ = fi45
        }
        switch v25 {
        case "foo bar":
            v48 := GenS1{F0: -104, F1: 2147483648}
            _ = v48
        case "\n\t":
            v4 = v36
        case "":
            v1 -= v1
        }
        fmt.Println(v1)
    } else {
        v3 += v3
        for fi49 := 0; fi49 < 15; fi49++ {
            v3 = (v3 + v3)
            _ = fi49
        }
    }
    for fi50 := 0; fi50 < 15; fi50++ {
        if false {
            fmt.Println(v36, v3, v3)
            var v51 uint32 = 593
            fmt.Println(v51, v51)
            _ = v51
        }
        v52 := GenS4{F0: v25}
        _ = fi50
        _ = v52
    }
    if true {
        v53 := GenS3{F0: 127, F1: 597, F2: 60}
        var v54 float32 = (float32(v42) + float32(((-2.25 + ((float64(1e10) / (v42 + (((v42 - v42) * ((float64(1e10) - (3.14159 - (1e10 / (v42 * (v42 / 0.1))))) * float64(0.0))) * ((v42 * 1e10) + v42)))) + (v42 / (-2.25 * v42)))) + 1.5)))
        var v55 uint = 42
        _ = v53
        _ = v54
        _ = v55
    } else {
        for fi56 := 0; fi56 < 20; fi56++ {
            v4 -= (v4 >> ((1) & 63))
            _ = fi56
        }
    }
    switch ("foo bar" + "hello") {
    case "\n\t":
        v57 := []int{32768}
        _ = v57
    case "":
        v36 = 0
    }
    v58 := GenS3{F0: 551, F1: int32(v3), F2: -128}
    var v59 uint64 = 0
    v3 += 1
    v60 := GenS5{F0: v59}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v4
    _ = v42
    _ = v58
    _ = v59
    _ = v60
}
