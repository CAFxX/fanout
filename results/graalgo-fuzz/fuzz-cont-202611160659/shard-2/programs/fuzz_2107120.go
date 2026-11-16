package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 string; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int32 }

func main() {
    var v1 uint16 = 32767
    v2 := []uint32{1, 2147483647, 0, 42}
    v3 := make(map[uint32]uint32)
    var v4 uint16 = (v1 >> ((2) & 63))
    var v5 bool = true
    v6 := GenS1{F0: -1}
    for fi7 := 0; fi7 < 9; fi7++ {
        v5 = v5
        switch "\n\t" {
        case "a":
            v8 := GenS1{F0: -1}
            _ = v8
        case "foo bar":
            if false {
                var v9 int16 = 907
                v9 = -32768
                var v10 string = ("foo bar" + ("\n\t" + "foo bar"))
                _ = v10
                _ = v9
            } else {
                if v5 {
                    v4 = (v4 + v4)
                    for fi11 := 0; fi11 < 10; fi11++ {
                        v4 = 29
                        v5 = (!(!(596 != 2147483648)))
                        switch (fi11 % ((fi7) | 1)) {
                        case -832:
                            v2 = append(v2, 65535)
                        case 890:
                            v12 := GenS2{F0: 18446744073709551615}
                            _ = v12
                        }
                        _ = fi11
                    }
                    var v13 string = ""
                    _ = v13
                }
                if v5 {
                    var v14 int = -772
                    v14 = fi7
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 7; fi15++ {
                        var v16 string = "a"
                        _ = fi15
                        _ = v16
                    }
                }
            }
        case "hello":
            v17 := GenS1{F0: -608}
            _ = v17
        }
        _ = fi7
    }
    switch "\n\t" {
    case "a":
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := v3[715]
            _ = fi18
            _ = v19
        }
    case "hello":
        v4 += v4
    default:
        v20 := GenS3{F0: 70, F1: "hello", F2: (v1 << ((683) & 63))}
        _ = v20
    }
    v4 -= v1
    _ = len(v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 = v4
    v21 := GenS2{F0: 18446744073709551615}
    switch 127 {
    case 42:
        for fi22 := 0; fi22 < 18; fi22++ {
            v5 = true
            _ = fi22
        }
    case -253:
        fmt.Println(v4, v4, v4)
    case -1:
        _ = len(v3)
    }
    var v23 float64 = 0.0
    var v24 string = "\n\t"
    delete(v3, 0)
    switch ((v24 + v24) + v24) {
    case "\n\t":
        v25 := GenS4{F0: 42, F1: -819}
        _ = v25
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        delete(v3, 4294967295)
        v3[685] = 42
        for fi27 := 0; fi27 < 1; fi27++ {
            switch (fi27 ^ (fi26 - -152)) {
            case 32767:
                if v5 {
                    if ((255 > 1) && (v24 == (v24 + (v24 + (v24 + v24))))) {
                        delete(v3, 426)
                        var v28 int16 = -1
                        if (true || (v23 >= float64(1.5))) {
                            if (2147483647 >= 42) {
                                var v29 string = v24
                                var v30 float32 = ((1.5 + (0.0 * (3.14159 * -2.25))) * (((0.1 - 1e10) - 0.1) * 1.5))
                                _ = v29
                                _ = v30
                            } else {
                                v31 := GenS3{F0: 70, F1: v24, F2: (^(v1))}
                                v5 = v5
                                _ = v31
                            }
                            v32 := GenS3{F0: 127, F1: "a", F2: 255}
                            _ = v32
                        }
                        _ = v28
                    } else {
                        var v33 float64 = v23
                        v1 -= 762
                        _ = v33
                    }
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                }
            }
            v34 := v3[1]
            var v35 uint16 = 256
            switch v24 {
            case "hello":
                var v36 int32 = -2147483648
                _ = v36
            case "\n\t":
                v24 = ((v24 + v24) + (v24 + v24))
            }
            v37 := GenS4{F0: 0, F1: 42}
            var v38 float32 = 0.0
            _ = fi27
            _ = v34
            _ = v35
            _ = v37
            _ = v38
        }
        _ = fi26
    }
    if len(v2) > 0 {
        v2[0] = 2
    }
    var v39 float64 = (1e10 * ((0.0 + -2.25) + 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
