package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint16 }

func main() {
    var v1 int32 = 0
    v2 := []int{-2147483648, 42}
    v3 := []int{32768}
    v1 = v1
    if len(v3) > 0 {
        v3[0] = -571
    }
    v1 += v1
    var v4 bool = (!false)
    var v5 bool = (0.0 < (-2.25 - (0.1 - 1e10)))
    var v6 int16 = -306
    var v7 uint16 = 2
    if len(v2) > 0 {
        v2[0] = 0
    }
    if v4 {
        if v4 {
            switch "hello" {
            case "a":
                for fi8 := 0; fi8 < 17; fi8++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v9 := GenS1{F0: fi8, F1: 0, F2: (v7 >> ((4294967295) & 63))}
                    _ = fi8
                    _ = v9
                }
            case "":
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 float32 = 1.5
                    switch "hello" {
                    case "foo bar":
                        var v12 bool = (v6 == 0)
                        _ = v12
                    case "\n\t":
                        for fi13 := 0; fi13 < 17; fi13++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v14 bool = v4
                            if ((!v14) && true) {
                                if ((v7 / ((0) | 1)) == v7) {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v15 bool = (!v4)
                                    _ = v15
                                } else {
                                    var v16 int64 = -1
                                    _ = v16
                                }
                                fmt.Println(v14, v14, v6)
                            } else {
                                v17 := GenS1{F0: 42, F1: 4294967295, F2: v7}
                                _ = v17
                            }
                            _ = fi13
                            _ = v14
                        }
                    }
                    _ = fi10
                    _ = v11
                }
            }
        }
    } else {
        if (v4 && v5) {
            if (!(v5 || (v1 <= 2147483647))) {
                for fi18 := 0; fi18 < 15; fi18++ {
                    fmt.Println(v6)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi18
                }
                v6 -= (v6 << ((18446744073709551615) & 63))
            } else {
                if (-1 > -1) {
                    v4 = v5
                    var v19 uint = 255
                    _ = v19
                }
                fmt.Println(v7, v4, v7)
            }
            var v20 int = -494
            _ = v20
        }
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 uint = 911
            fmt.Println(v6, v7)
            _ = fi21
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
