package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int32 }

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[string]string)
    v3 := make(map[uint32]uint32)
    var v4 string = ("hello" + "foo bar")
    v5 := GenS1{F0: -857, F1: 9223372036854775807}
    var v6 uint64 = 589
    v6 += v6
    if false {
        fmt.Println(v6)
        v6 -= v6
        if ("" >= v4) {
            v7 := GenS1{F0: -240, F1: -9223372036854775808}
            var v8 int64 = 42
            _ = v7
            _ = v8
        } else {
            fmt.Println(v6, v6)
        }
    } else {
        v4 = (((v4 + v4) + (v4 + "hello")) + v4)
    }
    v4 = (v4 + v4)
    if true {
        if (!false) {
            v4 = v4
            switch (v4 + v4) {
            case "foo bar":
                if (18446744073709551615 == 828) {
                    v4 = v4
                    if (!(1.5 < 0.1)) {
                        if true {
                            var v9 uint8 = 2
                            v9 = 255
                            for fi10 := 0; fi10 < 18; fi10++ {
                                var v11 int64 = -9223372036854775808
                                fmt.Println(v4)
                                v6 += 722
                                _ = fi10
                                _ = v11
                            }
                            _ = v9
                        } else {
                            v4 = (v4 + "foo bar")
                            var v12 uint = 127
                            _ = v12
                        }
                        if (42 >= 4) {
                            switch 0 {
                            case 32768:
                                v13 := GenS2{F0: 0, F1: 694, F2: 42}
                                _ = v13
                            }
                            v6 = v6
                        }
                        var v14 int8 = 0
                        _ = v14
                    }
                    switch 2147483647 {
                    case 127:
                        v4 = v4
                    }
                }
            }
        }
        var v15 int16 = -1
        for fi16 := 0; fi16 < 7; fi16++ {
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 uint32 = 1
                if true {
                    fmt.Println(fi16, v6)
                    var v19 int32 = 390
                    _ = v19
                }
                var v20 float32 = float32(0.1)
                _ = fi17
                _ = v18
                _ = v20
            }
            _ = fi16
        }
        _ = v15
    } else {
        fmt.Println(v6)
        fmt.Println(v4, v4)
        if len(v1) > 0 {
            v1[0] = v4
        }
    }
    v21 := GenS2{F0: 127, F1: 4294967295, F2: 42}
    var v22 int8 = -36
    if false {
        if (true && false) {
            var v23 uint16 = uint16(v6)
            switch -76 {
            case 285:
                if false {
                    for fi24 := 0; fi24 < 8; fi24++ {
                        v1 = append(v1, ("a" + v4))
                        _ = fi24
                    }
                }
            }
            _ = v23
        }
    }
    var v25 bool = (836 < 1)
    v1 = append(v1, ("hello" + ""))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
