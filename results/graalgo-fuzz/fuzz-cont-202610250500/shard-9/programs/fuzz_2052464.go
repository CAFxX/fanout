package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 int = 32768
    var v2 int = 328
    var v3 int8 = 127
    v4 := make(map[int]uint32)
    v5 := []int{(v2 ^ v1), 128, v1}
    v6 := GenS1{F0: true}
    _ = len(v4)
    var v7 uint = 677
    v1 = (v1 - 79)
    if true {
        if ((617 >= 65535) && false) {
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 int16 = 1
                if true {
                    v10 := v4[(v2 % ((v1) | 1))]
                    if (0 != 12) {
                        if (-2147483648 != 519) {
                            for fi11 := 0; fi11 < 13; fi11++ {
                                v12 := GenS2{F0: ("\n\t" + "foo bar")}
                                v5 = append(v5, (v1 - -2147483648))
                                switch (v2 * (v2 + fi11)) {
                                case -1:
                                    var v13 float32 = (-2.25 + 3.14159)
                                    _ = v13
                                case 32768:
                                    var v14 uint32 = (^(v10))
                                    _ = v14
                                }
                                v15 := v4[v2]
                                if len(v5) > 0 {
                                    v5[0] = (-(fi11))
                                }
                                v15 += v10
                                if (!true) {
                                    v4[v2] = 170
                                    fmt.Println(fi8)
                                } else {
                                    v15 += 42
                                }
                                _ = fi11
                                _ = v12
                                _ = v15
                            }
                        }
                    } else {
                        v16 := GenS3{F0: 2147483648, F1: false}
                        _ = v16
                    }
                    v2 += v1
                    _ = v10
                } else {
                    fmt.Println(v1)
                    for fi17 := 0; fi17 < 15; fi17++ {
                        if (true || false) {
                            if (84 >= (fi17 | (fi8 & (fi17 % (((^(fi8))) | 1))))) {
                                var v18 bool = false
                                if v18 {
                                    var v19 uint64 = 2
                                    _ = v19
                                } else {
                                    fmt.Println(v18)
                                }
                                v20 := GenS1{F0: (v18 || v18)}
                                _ = v18
                                _ = v20
                            }
                            fmt.Println(v3, v9, v7)
                            var v21 float32 = 1.5
                            if (724 == 127) {
                                var v22 uint64 = 4294967295
                                _ = v22
                            }
                            _ = v21
                        }
                        _ = fi17
                    }
                    if (42 < 1) {
                        if (-2147483648 != 2147483647) {
                            v4[(v2 - v2)] = 0
                            var v23 int8 = v3
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 17; fi24++ {
                            fmt.Println(v9, v9)
                            v7 += (v7 >> ((v7) & 63))
                            _ = len(v4)
                            _ = fi24
                        }
                        var v25 uint16 = 792
                        _ = v25
                    } else {
                        var v26 int16 = (v9 % ((int16(v7)) | 1))
                        var v27 uint8 = 19
                        _ = v26
                        _ = v27
                    }
                }
                _ = fi8
                _ = v9
            }
            var v28 uint = v7
            var v29 uint8 = 42
            _ = v28
            _ = v29
        }
        switch (v1 & (v2 + v2)) {
        case 1:
            if (false && (!(!true))) {
                v30 := GenS1{F0: true}
                var v31 int32 = 875
                for fi32 := 0; fi32 < 19; fi32++ {
                    v33 := v4[v1]
                    var v34 uint32 = v33
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                _ = v30
                _ = v31
            } else {
                v5 = append(v5, v2)
                v2 = v2
            }
        }
    } else {
        var v35 float32 = -2.25
        v36 := GenS4{F0: 173}
        for fi37 := 0; fi37 < 13; fi37++ {
            v3 = v3
            v1 -= 127
            _ = fi37
        }
        _ = v35
        _ = v36
    }
    var v38 int8 = v3
    var v39 bool = false
    delete(v4, v2)
    v7 = 65535
    switch (("" + "foo bar") + "\n\t") {
    case "hello":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "a":
        fmt.Println(v3, v38)
    case "":
        if len(v5) > 0 {
            v5[0] = (v1 | 1)
        }
    }
    v40 := GenS2{F0: ""}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, 128)
    v41 := GenS4{F0: 2147483648}
    switch -1 {
    case -264:
        var v42 uint8 = 2
        _ = v42
    case -210:
        fmt.Println(v7, v38)
    case 32767:
        fmt.Println(v2)
    default:
        v7 = v7
    }
    v38 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
