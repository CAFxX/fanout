package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int64 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 1
    var v3 uint32 = 62
    var v4 int = 0
    v5 := []int{(v4 * v4), v4}
    var v6 int16 = int16(v3)
    delete(v1, v4)
    if (true && false) {
        v7 := v1[(v4 >> ((212) & 63))]
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 string = "hello"
            _ = fi8
            _ = v9
        }
        v1[(v4 ^ v4)] = true
        _ = v7
    }
    switch "hello" {
    case "foo bar":
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 int = (v4 | v4)
            switch -9223372036854775808 {
            case 2147483647:
                fmt.Println(v2, v11, v2)
            case 9223372036854775807:
                v12 := GenS1{F0: (^(v3)), F1: 900, F2: 585}
                _ = v12
            case 128:
                v13 := v1[4294967295]
                _ = v13
            }
            var v14 uint32 = v3
            _ = fi10
            _ = v11
            _ = v14
        }
    case "a":
        var v15 uint = 926
        _ = v15
    case "":
        var v16 int16 = (v6 << ((42) & 63))
        _ = v16
    default:
        v17 := GenS2{F0: (v6 / ((int16(v4)) | 1))}
        _ = v17
    }
    v18 := GenS2{F0: v6}
    delete(v1, (-(v4)))
    var v19 string = "\n\t"
    for fi20 := 0; fi20 < 13; fi20++ {
        fmt.Println(v3, v6)
        _ = fi20
    }
    var v21 int64 = -1
    _ = len(v1)
    v21 = -73
    var v22 string = (v19 + "a")
    if false {
        fmt.Println(v19, v3)
    }
    if false {
        if true {
            var v23 uint32 = v3
            _ = v23
        } else {
            v22 = v19
            for fi24 := 0; fi24 < 19; fi24++ {
                for fi25 := 0; fi25 < 20; fi25++ {
                    if true {
                        var v26 int8 = -128
                        _ = v26
                    } else {
                        var v27 uint8 = 127
                        if true {
                            var v28 int32 = -2147483648
                            _ = v28
                        }
                        var v29 int = 32768
                        _ = v27
                        _ = v29
                    }
                    v30 := GenS1{F0: (v3 & v3), F1: 2147483647, F2: (v21 - v21)}
                    if (398 > 18446744073709551615) {
                        switch v22 {
                        case "a":
                            v5 = append(v5, int(v2))
                        }
                        _ = len(v5)
                    } else {
                        if false {
                            v31 := GenS1{F0: 2, F1: -45, F2: int64(fi24)}
                            fmt.Println(v21, fi25)
                            _ = v31
                        } else {
                            v2 = (v2 >> ((826) & 63))
                            fmt.Println(v22, v3, fi24)
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            switch (fi24 << ((65535) & 63)) {
                            case -1:
                                var v32 uint64 = (v2 << ((18446744073709551615) & 63))
                                _ = v32
                            case 42:
                                if false {
                                    var v33 bool = (510 < 301)
                                    _ = v33
                                } else {
                                    var v34 int = fi24
                                    _ = v34
                                }
                            }
                        }
                    }
                    _ = fi25
                    _ = v30
                }
                _ = fi24
            }
        }
        v35 := GenS1{F0: v3, F1: 337, F2: (v21 % (((v21 >> ((447) & 63))) | 1))}
        v21 = v21
        _ = v35
    } else {
        for fi36 := 0; fi36 < 2; fi36++ {
            var v37 uint16 = 255
            if true {
                fmt.Println(v19, v2, v6)
            }
            _ = fi36
            _ = v37
        }
        v38 := GenS2{F0: 32767}
        v19 = ("foo bar" + v22)
        _ = v38
    }
    if (-1 != -101) {
        if (!(!true)) {
            v3 = v3
        }
        switch (v4 - v4) {
        case 128:
            fmt.Println(v22, v4, v21)
        case 32768:
            for fi39 := 0; fi39 < 8; fi39++ {
                var v40 float64 = 1e10
                var v41 bool = true
                var v42 uint64 = (v2 - (v2 / (((v2 & (^(v2)))) | 1)))
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
            }
        }
        var v43 float32 = 0.0
        _ = v43
    } else {
        var v44 bool = false
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
