package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int32) int64 {
    v4 := GenS1{F0: 65535}
    var v5 string = "foo bar"
    if (!(-539 <= 1)) {
        if (-128 >= -62) {
            if true {
                var v6 int = 32767
                _ = v6
            } else {
                switch 127 {
                case 32767:
                    if true {
                        v7 := GenS1{F0: 42}
                        var v8 uint64 = 18446744073709551615
                        if true {
                            p0 = (p0 >> ((939) & 63))
                            fmt.Println(v8, v8, p0)
                            if (1 <= 128) {
                                var v9 uint = 0
                                _ = v9
                            }
                        }
                        _ = v7
                        _ = v8
                    } else {
                        v10 := GenS1{F0: 2}
                        if false {
                            if (908 > 2147483647) {
                                var v11 bool = false
                                _ = v11
                            } else {
                                v5 = v5
                                switch 42 {
                                case 2147483647:
                                    fmt.Println(p0, v5)
                                case 128:
                                    var v12 int8 = 29
                                    _ = v12
                                }
                            }
                            if false {
                                p0 -= p0
                                p0 += 42
                                var v13 int32 = p0
                                _ = v13
                            }
                        }
                        v5 = "hello"
                        _ = v10
                    }
                }
            }
        }
    } else {
        v14 := GenS1{F0: 42}
        v15 := GenS1{F0: 2147483647}
        for fi16 := 0; fi16 < 6; fi16++ {
            p0 = p0
            _ = fi16
        }
        _ = v14
        _ = v15
    }
    var v17 uint8 = 0
    _ = p0
    _ = v17
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 int32 = -831
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    for fi18 := 0; fi18 < 10; fi18++ {
        v19 := v2[""]
        _ = fi18
        _ = v19
    }
    delete(v2, "hello")
    v1 = v1
    var v20 uint32 = 2
    switch 143 {
    case -521:
        v20 += (^(v20))
    case -2147483648:
        v1 += v1
    case 32768:
        if (float32(0.1) == (float32(-2.25) - 0.0)) {
            v21 := GenS1{F0: v20}
            _ = v21
        } else {
            switch -2147483648 {
            case -1:
                v20 = v20
            case 42:
                var v22 uint16 = 42
                _ = v22
            case 769:
                if (18446744073709551615 < 1) {
                    var v23 uint16 = 687
                    _ = v23
                } else {
                    v1 = (v1 / ((v1) | 1))
                    delete(v2, "foo bar")
                }
            }
            if false {
                v1 = 0
            }
        }
    }
    v20 = (^(v20))
    var v24 int8 = int8(v1)
    delete(v3, v20)
    v24 -= (v24 & v24)
    var v25 bool = (!(v20 > v20))
    v26 := GenS2{F0: (v24 & v24)}
    var v27 int64 = 0
    v2["a"] = v20
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
