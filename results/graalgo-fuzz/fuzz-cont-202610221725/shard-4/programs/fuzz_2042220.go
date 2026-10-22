package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0(p0 int) bool {
    v4 := []int{}
    for fi5 := 0; fi5 < 10; fi5++ {
        p0 = p0
        var v6 uint64 = 274
        p0 = 32768
        _ = fi5
        _ = v6
    }
    if false {
        p0 = (p0 | (^(p0)))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    switch (p0 | p0) {
    case -674:
        var v7 string = "hello"
        _ = v7
    case 127:
        if ((true || ((p0 % ((2147483647) | 1)) == p0)) || (p0 >= 0)) {
            var v8 uint8 = 128
            var v9 int16 = -32768
            v10 := GenS1{F0: p0}
            _ = v10
            _ = v8
            _ = v9
        }
    default:
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 int64 = 42
            v12 -= (v12 - v12)
            var v13 int16 = 1000
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    _ = p0
    _ = v4
    return false
}

func main() {
    v1 := []uint32{916, 1, 42, 825}
    v2 := make(map[uint32]uint32)
    var v3 int64 = 9223372036854775807
    if (!true) {
        var v14 bool = (1e10 <= 1.5)
        fmt.Println(v3, v14)
        switch 4294967295 {
        case 0:
            fmt.Println(v3, v14)
        }
        _ = v14
    } else {
        var v15 uint8 = 1
        _ = v15
    }
    var v16 int64 = v3
    _ = len(v2)
    var v17 uint32 = 2147483648
    if false {
        v17 += 739
    }
    var v18 bool = false
    for fi19 := 0; fi19 < 2; fi19++ {
        fmt.Println(v18, v18, v17)
        if (((!v18) && (363 < v17)) || v18) {
            for fi20 := 0; fi20 < 10; fi20++ {
                v3 -= (v16 % ((v16) | 1))
                var v21 string = "hello"
                var v22 int8 = 0
                _ = fi20
                _ = v21
                _ = v22
            }
        } else {
            var v23 uint16 = 1
            var v24 int64 = (v16 | v16)
            fmt.Println(v17, v17, v17)
            _ = v23
            _ = v24
        }
        _ = fi19
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        var v26 uint = 65535
        v27 := GenS1{F0: 32767}
        delete(v2, v17)
        _ = fi25
        _ = v26
        _ = v27
    }
    switch 0 {
    case 0:
        v3 = 9223372036854775807
    }
    v28 := GenS1{F0: 32767}
    var v29 uint64 = 4294967295
    switch 1 {
    case 0:
        v29 = 4294967295
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, v17)
    var v30 uint32 = v17
    v31 := GenS1{F0: 0}
    if false {
        if v18 {
            for fi32 := 0; fi32 < 9; fi32++ {
                v30 = (v17 & 2147483648)
                _ = fi32
            }
            if false {
                var v33 uint8 = 127
                _ = v33
            }
        } else {
            for fi34 := 0; fi34 < 9; fi34++ {
                for fi35 := 0; fi35 < 15; fi35++ {
                    v36 := v2[v17]
                    v16 = 42
                    if false {
                        var v37 uint8 = 1
                        _ = v37
                    }
                    v29 += (v29 / ((uint64(v30)) | 1))
                    _ = fi35
                    _ = v36
                }
                var v38 int16 = -32768
                _ = fi34
                _ = v38
            }
            for fi39 := 0; fi39 < 2; fi39++ {
                switch 9223372036854775807 {
                case 32767:
                    v40 := GenS1{F0: (fi39 - (fi39 * fi39))}
                    _ = v40
                case 2147483647:
                    for fi41 := 0; fi41 < 1; fi41++ {
                        switch fi41 {
                        case 42:
                            v42 := GenS2{F0: (42 == 377), F1: v30}
                            _ = v42
                        case -1:
                            var v43 uint32 = v17
                            _ = v43
                        case -2147483648:
                            var v44 string = "\n\t"
                            _ = v44
                        }
                        _ = fi41
                    }
                }
                v16 = 0
                v45 := GenS1{F0: (fi39 << ((18446744073709551615) & 63))}
                _ = fi39
                _ = v45
            }
            var v46 int16 = 42
            _ = v46
        }
        v17 = 0
        v16 -= v16
    } else {
        var v47 uint32 = v30
        for fi48 := 0; fi48 < 4; fi48++ {
            for fi49 := 0; fi49 < 7; fi49++ {
                fmt.Println(v18, fi48, v3)
                _ = fi49
            }
            v17 = 0
            _ = fi48
        }
        v50 := GenS2{F0: ((!v18) || (true && (100 <= 100))), F1: v30}
        _ = v47
        _ = v50
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
