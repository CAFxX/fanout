package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() int16 {
    var v4 uint16 = 42
    v5 := GenS1{F0: 2147483647}
    v6 := []int{-536, 9223372036854775807, -965}
    if len(v6) > 0 {
        v6[0] = 649
    }
    _ = v4
    _ = v5
    _ = v6
    return -32768
}

func main() {
    var v1 int64 = -1
    v2 := []int{}
    var v3 uint64 = 1
    for fi7 := 0; fi7 < 17; fi7++ {
        v2 = append(v2, fi7)
        _ = fi7
    }
    if len(v2) > 0 {
        v2[0] = -2147483648
    }
    if false {
        var v8 bool = (!(true && true))
        _ = v8
    }
    var v9 int = 989
    if (true || (!true)) {
        if false {
            fmt.Println(v9, v9)
            for fi10 := 0; fi10 < 9; fi10++ {
                fmt.Println(v9, v1, v1)
                v2 = append(v2, (v9 * v9))
                _ = fi10
            }
        }
        switch "hello" {
        case "\n\t":
            switch ("foo bar" + "a") {
            case "hello":
                fmt.Println(v9)
            case "a":
                v9 = int(v1)
            }
        case "":
            _ = len(v2)
        }
        var v11 uint16 = 1
        _ = v11
    } else {
        var v12 uint = 2
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 2; fi13++ {
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 int16 = -32768
                if false {
                    if false {
                        fmt.Println(v1)
                    }
                    var v16 float32 = float32(0.1)
                    _ = v16
                } else {
                    v9 -= fi13
                }
                _ = fi14
                _ = v15
            }
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 int16 = 604
                if true {
                    v9 += 1
                    var v19 float32 = 0.1
                    _ = v19
                }
                v18 -= -32768
                var v20 uint32 = 1
                var v21 string = ""
                _ = fi17
                _ = v18
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 16; fi22++ {
                var v23 uint8 = 42
                _ = fi22
                _ = v23
            }
            _ = fi13
        }
        if (!(2147483647 < 442)) {
            v1 += (^(v1))
            if (0.1 <= 1e10) {
                var v24 uint = 4294967295
                _ = v24
            }
        }
        _ = len(v2)
    }
    var v25 uint32 = 419
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 uint = 127
        _ = fi26
        _ = v27
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v28 int = v9
    fmt.Println(v9)
    var v29 string = ((("a" + "hello") + "") + "foo bar")
    switch 2147483647 {
    case -1:
        switch int(v25) {
        case -1:
            switch v9 {
            case -398:
                if len(v2) > 0 {
                    v2[0] = v28
                }
            case -2147483648:
                switch "" {
                case "foo bar":
                    switch "hello" {
                    case "a":
                        var v30 int8 = int8(v3)
                        _ = v30
                    case "foo bar":
                        v31 := GenS1{F0: -236}
                        _ = v31
                    case "":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                case "\n\t":
                    var v32 string = "a"
                    _ = v32
                case "a":
                    for fi33 := 0; fi33 < 11; fi33++ {
                        var v34 int = (v28 * (fi33 * (-(v9))))
                        if len(v2) > 0 {
                            v2[0] = (v34 >> ((2) & 63))
                        }
                        _ = fi33
                        _ = v34
                    }
                }
            case -1:
                v29 = "foo bar"
            }
        }
    case 2147483647:
        var v35 int64 = v1
        _ = v35
    case 445:
        v36 := GenS1{F0: 1}
        _ = v36
    }
    var v37 int = v9
    _ = len(v2)
    for fi38 := 0; fi38 < 16; fi38++ {
        for fi39 := 0; fi39 < 3; fi39++ {
            if (0.1 > (float32(float64(0.1)) / ((3.14159 - 0.0) * 3.14159))) {
                var v40 uint8 = 1
                v2 = append(v2, (v9 | fi39))
                _ = v40
            } else {
                var v41 int = v9
                _ = v41
            }
            _ = fi39
        }
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v9
}
