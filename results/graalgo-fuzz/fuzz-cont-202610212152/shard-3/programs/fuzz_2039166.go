package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int64 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint32 }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 10; fi4++ {
        switch ("\n\t" + "") {
        case "a":
            var v5 float64 = 1e10
            _ = v5
        }
        v6 := GenS1{F0: -332, F1: 100, F2: 42}
        var v7 int64 = -9223372036854775808
        _ = fi4
        _ = v6
        _ = v7
    }
    v8 := []int{}
    var v9 uint = 65535
    _ = v8
    _ = v9
    return 93
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = ("a" + "")
    v3 := []string{(v2 + ((v2 + "a") + (v2 + v2))), "\n\t", v2}
    v10 := GenS1{F0: -585, F1: -22, F2: 42}
    v11 := v1[0]
    v12 := GenS2{F0: 629, F1: 127}
    v13 := GenS1{F0: 361, F1: -111, F2: 42}
    v14 := GenS3{F0: 451, F1: 0}
    v15 := GenS1{F0: 31, F1: -1, F2: -9223372036854775808}
    v11 = 0
    var v16 int16 = 956
    fmt.Println(v2)
    if (false && true) {
        var v17 string = ""
        if (2 > 2) {
            v17 = v2
        }
        v18 := GenS3{F0: 476, F1: -128}
        _ = v17
        _ = v18
    } else {
        if false {
            v16 += (v16 >> ((730) & 63))
        } else {
            var v19 int64 = -239
            _ = v19
        }
        for fi20 := 0; fi20 < 19; fi20++ {
            v11 = (fi20 & (v11 & (^(v11))))
            v2 = ("hello" + (v2 + ("hello" + ((v2 + "\n\t") + ""))))
            var v21 uint16 = 255
            _ = fi20
            _ = v21
        }
    }
    var v22 int32 = -380
    if true {
        var v23 uint32 = 2147483647
        _ = v23
    } else {
        var v24 uint = 293
        v25 := GenS1{F0: 42, F1: -40, F2: -9223372036854775808}
        _ = v24
        _ = v25
    }
    v26 := GenS4{F0: 42}
    fmt.Println(v11, v16, v2)
    if (356 < 65535) {
        v27 := GenS4{F0: 2147483648}
        if false {
            v11 -= v11
            if false {
                if true {
                    v11 += (-(v11))
                    if (42 != 2) {
                        for fi28 := 0; fi28 < 17; fi28++ {
                            if (32767 > 1) {
                                for fi29 := 0; fi29 < 4; fi29++ {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v30 uint32 = 1
                                    _ = len(v3)
                                    _ = fi29
                                    _ = v30
                                }
                            } else {
                                var v31 int = 0
                                var v32 int32 = (v22 * (v22 * -777))
                                fmt.Println(v11, fi28, fi28)
                                _ = v31
                                _ = v32
                            }
                            _ = len(v1)
                            var v33 uint32 = 158
                            _ = fi28
                            _ = v33
                        }
                    }
                    switch 210 {
                    case 128:
                        if false {
                            var v34 int16 = 32767
                            fmt.Println(v34, v34, v22)
                            _ = v34
                        }
                    }
                }
                var v35 string = "\n\t"
                _ = v35
            }
            v36 := GenS2{F0: 32767, F1: 127}
            _ = v36
        }
        var v37 uint8 = genFn0()
        _ = v27
        _ = v37
    }
    switch (v2 + ((v2 + (("\n\t" + ("hello" + "hello")) + "\n\t")) + v2)) {
    case "":
        for fi38 := 0; fi38 < 4; fi38++ {
            v2 = ""
            _ = fi38
        }
    }
    for fi39 := 0; fi39 < 3; fi39++ {
        v40 := GenS2{F0: v16, F1: genFn0()}
        switch (v2 + v2) {
        case "hello":
            switch v11 {
            case -322:
                v41 := GenS1{F0: 0, F1: 127, F2: 641}
                _ = v41
            }
        }
        _ = fi39
        _ = v40
    }
    v42 := GenS4{F0: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v42
}
