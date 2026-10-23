package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []int{9223372036854775807, 42, -488}
    var v2 string = ""
    var v3 int64 = 0
    var v4 int32 = 0
    switch v2 {
    case "":
        var v5 float32 = 0.1
        _ = v5
    }
    var v6 int8 = 0
    for fi7 := 0; fi7 < 19; fi7++ {
        v6 += v6
        _ = fi7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        if false {
            fmt.Println(v2)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v9 int16 = 1000
        _ = fi8
        _ = v9
    }
    switch 4294967295 {
    case -9223372036854775808:
        switch "\n\t" {
        case "a":
            var v10 uint16 = 893
            _ = v10
        case "foo bar":
            var v11 int8 = int8(v3)
            _ = v11
        case "\n\t":
            for fi12 := 0; fi12 < 10; fi12++ {
                v2 = v2
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 uint32 = 0
                    var v15 float32 = 0.1
                    v16 := GenS1{F0: 186}
                    var v17 uint16 = 527
                    var v18 uint64 = 2
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = fi12
            }
        default:
            switch -698 {
            case 32768:
                var v19 int8 = v6
                _ = v19
            }
        }
    case 32767:
        if true {
            v20 := GenS1{F0: -32768}
            var v21 uint64 = 424
            for fi22 := 0; fi22 < 14; fi22++ {
                if true {
                    for fi23 := 0; fi23 < 3; fi23++ {
                        v2 = v2
                        if ((v3 * (v3 >> ((42) & 63))) <= (v3 << ((65535) & 63))) {
                            if (127 > 101) {
                                v2 = "a"
                                _ = len(v1)
                            } else {
                                v24 := GenS1{F0: 42}
                                v2 = ("\n\t" + v2)
                                fmt.Println(v21, fi22)
                                _ = v24
                            }
                        }
                        fmt.Println(v6, v4)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi23
                    }
                    _ = len(v1)
                }
                _ = fi22
            }
            _ = v20
            _ = v21
        }
    default:
        v6 -= (v6 - 100)
    }
    var v25 int8 = (v6 >> ((4294967295) & 63))
    var v26 uint16 = 42
    switch v2 {
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = -422
        }
    case "\n\t":
        var v27 float32 = float32(3.14159)
        _ = v27
    case "":
        for fi28 := 0; fi28 < 14; fi28++ {
            for fi29 := 0; fi29 < 1; fi29++ {
                v30 := GenS1{F0: 1000}
                fmt.Println(v26, v25, v4)
                _ = fi29
                _ = v30
            }
            if false {
                switch v2 {
                case "":
                    fmt.Println(v25, v26, v25)
                }
                v25 = -117
                if true {
                    if (true || (false && true)) {
                        for fi31 := 0; fi31 < 10; fi31++ {
                            var v32 uint8 = 0
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi31
                            _ = v32
                        }
                        var v33 float32 = (-2.25 + 1e10)
                        _ = v33
                    }
                } else {
                    v6 -= v6
                }
            }
            _ = fi28
        }
    }
    var v34 uint64 = 638
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v6
}
