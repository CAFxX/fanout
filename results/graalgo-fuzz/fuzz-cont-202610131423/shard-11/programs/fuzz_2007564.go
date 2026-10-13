package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 string = ""
    var v2 int32 = 1
    v3 := make(map[uint32]bool)
    var v4 int16 = 32767
    v5 := make(map[string]uint32)
    var v6 uint16 = 65535
    switch v1 {
    case "":
        for fi7 := 0; fi7 < 10; fi7++ {
            switch "\n\t" {
            case "foo bar":
                var v8 int16 = v4
                _ = v8
            case "hello":
                if false {
                    v9 := []int{(-(fi7)), (fi7 << ((65535) & 63))}
                    _ = v9
                }
            case "a":
                var v10 int8 = -1
                _ = v10
            }
            if true {
                var v11 int64 = 886
                _ = v11
            }
            _ = fi7
        }
    case "foo bar":
        v2 = (v2 << ((538) & 63))
    }
    if true {
        fmt.Println(v4)
        for fi12 := 0; fi12 < 6; fi12++ {
            v13 := GenS1{F0: 127}
            _ = fi12
            _ = v13
        }
        switch 829 {
        case 4294967295:
            v14 := GenS1{F0: 0}
            _ = v14
        }
    } else {
        var v15 uint = 758
        _ = v15
    }
    var v16 uint16 = (v6 ^ v6)
    if (46 <= 138) {
        var v17 int = -440
        var v18 float32 = 1.5
        _ = v17
        _ = v18
    }
    if true {
        _ = len(v3)
        for fi19 := 0; fi19 < 19; fi19++ {
            if true {
                for fi20 := 0; fi20 < 14; fi20++ {
                    if true {
                        v1 = v1
                    }
                    if true {
                        v21 := GenS1{F0: 127}
                        _ = v21
                    }
                    fmt.Println(fi19, v16, fi19)
                    _ = fi20
                }
                for fi22 := 0; fi22 < 12; fi22++ {
                    v16 = (v6 - (v6 / ((v6) | 1)))
                    var v23 uint8 = 2
                    var v24 int8 = 0
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
            switch fi19 {
            case -550:
                for fi25 := 0; fi25 < 12; fi25++ {
                    switch v1 {
                    case "foo bar":
                        v2 = v2
                    case "hello":
                        v26 := []int{fi19}
                        _ = v26
                    default:
                        v27 := GenS1{F0: -1}
                        _ = v27
                    }
                    fmt.Println(fi25, fi19, v6)
                    _ = fi25
                }
            case -2147483648:
                for fi28 := 0; fi28 < 1; fi28++ {
                    switch 32768 {
                    case 0:
                        if true {
                            if (!false) {
                                v4 = v4
                                var v29 int64 = 550
                                _ = v29
                            }
                            var v30 bool = true
                            switch v1 {
                            case "\n\t":
                                v6 = v16
                            case "foo bar":
                                delete(v5, v1)
                            case "hello":
                                var v31 int8 = 88
                                _ = v31
                            default:
                                v6 = (^(v16))
                            }
                            _ = v30
                        } else {
                            v1 = "hello"
                        }
                    }
                    v32 := v3[2147483648]
                    _ = fi28
                    _ = v32
                }
            case 2147483647:
                delete(v3, 713)
            default:
                var v33 int32 = (v2 | 42)
                _ = v33
            }
            _ = fi19
        }
        switch "foo bar" {
        case "":
            var v34 int16 = v4
            _ = v34
        default:
            v35 := []int{32767}
            _ = v35
        }
    } else {
        v16 -= v6
        v2 += (v2 << ((255) & 63))
        for fi36 := 0; fi36 < 8; fi36++ {
            var v37 int16 = 447
            v4 = v4
            _ = fi36
            _ = v37
        }
    }
    delete(v5, ("hello" + v1))
    for fi38 := 0; fi38 < 8; fi38++ {
        switch -318 {
        case 4294967295:
            v2 += v2
        case 389:
            fmt.Println(v2, v6, v2)
        default:
            if (!(((!false) && true) && false)) {
                v39 := GenS1{F0: -119}
                for fi40 := 0; fi40 < 19; fi40++ {
                    fmt.Println(fi40)
                    _ = fi40
                }
                v16 -= v16
                _ = v39
            }
        }
        _ = fi38
    }
    var v41 bool = false
    v5[v1] = 0
    switch -373 {
    case 32767:
        v1 = "a"
    case 1:
        var v42 uint32 = uint32(v16)
        _ = v42
    }
    for fi43 := 0; fi43 < 19; fi43++ {
        v1 = (v1 + "foo bar")
        _ = fi43
    }
    var v44 int8 = 100
    if v41 {
        v45 := []int{42, 42}
        _ = v45
    }
    switch (v1 + v1) {
    case "":
        v6 -= v16
    }
    v46 := GenS1{F0: (^(v44))}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v44
    _ = v46
    _ = v5
    _ = v6
}
