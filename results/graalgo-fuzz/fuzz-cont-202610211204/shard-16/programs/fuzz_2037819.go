package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint16; F1 int8; F2 bool }

func main() {
    var v1 string = (("" + ("a" + "foo bar")) + "foo bar")
    var v2 uint = 1
    v3 := []int{}
    fmt.Println(v2, v2)
    for fi4 := 0; fi4 < 13; fi4++ {
        for fi5 := 0; fi5 < 14; fi5++ {
            v2 += (^(v2))
            v6 := GenS1{F0: 42, F1: (577 > 4294967295)}
            v1 = v1
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    var v7 uint64 = 4294967295
    if (-1 <= 100) {
        v1 = (v1 + ((("foo bar" + v1) + v1) + v1))
        for fi8 := 0; fi8 < 3; fi8++ {
            fmt.Println(v1, fi8)
            _ = fi8
        }
    } else {
        switch v1 {
        case "hello":
            v1 = "hello"
        case "foo bar":
            var v9 uint64 = (v7 & 0)
            _ = v9
        case "a":
            v7 = (v7 << (((v2 >> ((v2) & 63))) & 63))
        default:
            fmt.Println(v7, v1)
        }
        v7 -= 772
    }
    fmt.Println(v1, v2)
    var v10 float64 = 0.0
    v11 := GenS1{F0: 629, F1: true}
    for fi12 := 0; fi12 < 6; fi12++ {
        switch 232 {
        case 0:
            var v13 uint8 = 128
            _ = v13
        case 1:
            switch ("hello" + "foo bar") {
            case "a":
                v1 = v1
            case "hello":
                var v14 int16 = 32767
                _ = v14
            default:
                v15 := GenS1{F0: 1, F1: true}
                _ = v15
            }
        default:
            v1 = "a"
        }
        _ = len(v3)
        _ = fi12
    }
    var v16 int8 = 42
    switch 725 {
    case 1:
        if false {
            var v17 uint32 = 2147483648
            v1 = (v1 + "foo bar")
            v18 := GenS1{F0: 0, F1: false}
            _ = v17
            _ = v18
        } else {
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 uint8 = 2
                _ = fi19
                _ = v20
            }
        }
    case 667:
        switch "hello" {
        case "hello":
            _ = len(v3)
        case "a":
            var v21 float32 = -2.25
            _ = v21
        case "foo bar":
            switch ((v1 + "hello") + v1) {
            case "\n\t":
                var v22 uint8 = 61
                _ = v22
            case "foo bar":
                if (!true) {
                    v23 := GenS2{F0: 65535, F1: v16, F2: true}
                    v16 += (v16 / ((42) | 1))
                    _ = v23
                } else {
                    v16 += v16
                    v16 += (v16 - (v16 * -1))
                }
            }
        }
    }
    if false {
        if false {
            v24 := GenS1{F0: -827, F1: false}
            _ = v24
        } else {
            for fi25 := 0; fi25 < 1; fi25++ {
                var v26 int8 = 117
                var v27 int16 = 1
                for fi28 := 0; fi28 < 2; fi28++ {
                    v3 = append(v3, (fi28 - fi25))
                    _ = fi28
                }
                _ = fi25
                _ = v26
                _ = v27
            }
            var v29 float32 = (0.1 * float32(float64(float32(v10))))
            v10 = v10
            _ = v29
        }
        fmt.Println(v2)
    } else {
        for fi30 := 0; fi30 < 19; fi30++ {
            v10 -= v10
            if false {
                if true {
                    switch (^(fi30)) {
                    case -1:
                        if (!(1 <= 801)) {
                            switch (fi30 >> ((v2) & 63)) {
                            case 929:
                                v31 := GenS2{F0: 255, F1: (v16 % (((^(v16))) | 1)), F2: (true || (!(fi30 != fi30)))}
                                _ = v31
                            }
                            fmt.Println(v10, fi30, fi30)
                        }
                    case -974:
                        var v32 uint = 65535
                        _ = v32
                    }
                } else {
                    v33 := GenS2{F0: 32767, F1: 42, F2: false}
                    _ = v33
                }
                fmt.Println(v2, v10)
            }
            _ = fi30
        }
        fmt.Println(v10, v7)
    }
    var v34 bool = true
    var v35 bool = v34
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v7
}
